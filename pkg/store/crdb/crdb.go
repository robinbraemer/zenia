package crdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Crdb struct {
	conn *pgx.Conn
}

func New(ctx context.Context, connString string) (*Crdb, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting: %w", err)
	}
	return &Crdb{conn: conn}, nil
}

// CloseContext closes the database connection.
// It is safe to call when already closed.
func (db *Crdb) CloseContext(ctx context.Context) error {
	return db.conn.Close(ctx)
}
