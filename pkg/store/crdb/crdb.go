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

func (db *Crdb) Save(ctx context.Context, tuple acl.RelationTuple) error {
	panic("implement me")
}

func (db *Crdb) GetNamespaces(ctx context.Context) ([]acl.Namespace, error) {
	panic("implement me")
}

const (
	// Namespaced relation tuple table
	// TODO beware of sql injection!
	relationTupleTable = `
CREATE TABLE %s_relation_tuples (
    shard_id VARCHAR,
    object_id VARCHAR,
    relation VARCHAR,
    user VARCHAR,
    commit_time TIMESTAMP,
    PRIMARY KEY (shard_id, object_id, relation, user, commit_time)
)
`
	insertNamespace = `
INSERT INTO namespace_config
(namespace,config,commit_time)
VALUES(?,?,?)
`
)

func (db *Crdb) SaveNamespace(ctx context.Context, namespace acl.Namespace) error {
	config, commit := new(bytes.Buffer), time.Now().UTC()
	err := json.NewEncoder(config).Encode(&namespace.NamespaceConfig)
	if err != nil {
		return fmt.Errorf("error encoding namespace config: %w", err)
	}

	err = crdbpgx.ExecuteTx(ctx, db.Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, insertNamespace, namespace.Name, config, commit)
		if err != nil {
			return fmt.Errorf("error inserting namespace: %w", err)
		}

		// Create relation table
		_, err = tx.Exec(ctx, fmt.Sprintf(relationTupleTable, namespaceTable(namespace.Name)))
		if err != nil {
			return fmt.Errorf("error creating namespaced relation tuple table: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error executing db transaction: %w", err)
	}
	return nil
}

func namespaceTable(namespace string) (table string) {
	return fmt.Sprintf("%s_relation_tuple", namespace)
}
