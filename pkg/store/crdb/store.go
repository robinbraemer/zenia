package crdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/robinbraemer/zenia/pkg/acl"
	"time"
)

type Crdb struct {
	*pgx.Conn
}

func New(ctx context.Context, connString string) (*Crdb, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting: %w", err)
	}
	return &Crdb{Conn: conn}, nil
}

// CloseContext closes the database connection.
// It is safe to call when already closed.
func (db *Crdb) CloseContext(ctx context.Context) error {
	if db == nil || db.Conn == nil {
		return nil
	}
	return db.Conn.Close(ctx)
}

func (db *Crdb) Exists(ctx context.Context, object acl.Object, relation string, userID string) (bool, error) {
	panic("implement me")
}

func (db *Crdb) UserSets(ctx context.Context, object acl.Object, relation string) ([]acl.UserSet, error) {
	panic("implement me")
}

const insertTuple = `
INSERT INTO %s
(shard_id, object_id, relation, "user", commit_time)
VALUES ($1,$2,$3,$4,$5)
`

func (db *Crdb) Save(ctx context.Context, tuple acl.RelationTuple) error {
	_, err := db.Exec(ctx, fmt.Sprintf(insertTuple, namespacedTupleTable(tuple.Object.Namespace)),
		db.newShardID(tuple.Object.ID),
		tuple.Object.ID,
		tuple.Relation,
		tuple.User.String(),
		time.Now().UTC(),
	)
	return err
}

// TODO based of acl.NamespaceStorageSettings settings.Sharding.ComputedIDExpression
func (db *Crdb) newShardID(objectID string) string {
	return objectID
}

// Get newest config for each namespace.
//
// Remember: Multiple configs can get inserted for the same namespace
// at different commit_time and are kept within a garbage collection
// window so ACL checks with older Zookies refer to older committed namespace configs.
const selectLatestNamespaceConfigs = `
SELECT c1.namespace AS namespace, c1.config AS config
FROM namespace_config c1
LEFT JOIN namespace_config c2
  ON (c1.namespace = c2.namespace AND c1.commit_time < c2.commit_time)
`

func (db *Crdb) GetNamespaces(ctx context.Context) (l []acl.Namespace, err error) {
	rows, err := db.Query(ctx, selectLatestNamespaceConfigs)
	if err != nil {
		return nil, fmt.Errorf("error query latest namespace configs: %w", err)
	}
	var (
		name   string
		config acl.NamespaceConfig
	)
	for rows.Next() {
		if err = rows.Scan(&name, &config); err != nil {
			return nil, fmt.Errorf("error scan row: %w", err)
		}
		l = append(l, acl.Namespace{
			Name:            name,
			NamespaceConfig: config,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}
	return l, nil
}

const (
	// Namespaced relation tuple table
	// TODO beware of sql injection!
	tupleTable = `
CREATE TABLE %s (
    shard_id VARCHAR,
    object_id VARCHAR,
    relation VARCHAR,
    "user" VARCHAR,
    commit_time TIMESTAMP,
    PRIMARY KEY (shard_id, object_id, relation, "user", commit_time)
)
`
	insertNamespace = `
INSERT INTO namespace_config
(namespace,config,commit_time)
VALUES($1,$2,$3)
`
)

func (db *Crdb) SaveNamespace(ctx context.Context, namespace acl.Namespace) error {
	config, commit := new(bytes.Buffer), time.Now().UTC()
	err := json.NewEncoder(config).Encode(&namespace.NamespaceConfig)
	if err != nil {
		return fmt.Errorf("error encoding namespace config: %w", err)
	}

	err = crdbpgx.ExecuteTx(ctx, db.Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		// Create relation table
		_, err = tx.Exec(ctx, fmt.Sprintf(tupleTable, namespacedTupleTable(namespace.Name)))
		if err != nil {
			return fmt.Errorf("error creating namespaced relation tuple table: %w", err)
		}

		// Insert namespace config
		_, err := tx.Exec(ctx, insertNamespace, namespace.Name, config, commit)
		if err != nil {
			return fmt.Errorf("error inserting namespace: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error executing db transaction: %w", err)
	}
	return nil
}

func namespacedTupleTable(namespace string) (table string) {
	return fmt.Sprintf("%s_relation_tuple", namespace)
}
