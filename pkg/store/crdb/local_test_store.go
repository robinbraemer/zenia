package crdb

import (
	"context"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"io/ioutil"
)

func NewTestStore(ctx context.Context, migrateSqlFiles ...string) (store *Crdb, err error) {
	// Start local cockroachdb
	localNode, err := testserver.NewTestServer()
	if err != nil {
		return nil, fmt.Errorf("error starting local server: %v", err)
	}
	go func() { <-ctx.Done(); localNode.Stop() }()

	// Create db client
	store, err = New(ctx, localNode.PGURL().String())
	if err != nil {
		return nil, fmt.Errorf("error new client: %v", err)
	}
	go func() { <-ctx.Done(); store.CloseContext(ctx) }()

	// Migrate db
	if err := migrations(ctx, store, migrateSqlFiles...); err != nil {
		return nil, fmt.Errorf("error migrate: %v", err)
	}
	return store, nil
}

func migrations(ctx context.Context, store *Crdb, migrateSqlFiles ...string) error {
	for _, file := range migrateSqlFiles {
		if err := migrate(ctx, store, file); err != nil {
			return fmt.Errorf("error migrate file %s: %v", file, err)
		}
	}
	return nil
}

func migrate(ctx context.Context, store *Crdb, sqlFile string) error {
	sql, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("error reading migration file %s: %v", sqlFile, err)
	}
	_, err = store.Exec(ctx, string(sql))
	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}
	return nil
}
