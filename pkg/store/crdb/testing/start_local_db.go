package testing

import (
	"context"
	"fmt"
	"github.com/robinbraemer/zenia/pkg/store/crdb"
	"os/exec"
)

func StartLocalSingleNode(ctx context.Context) (*crdb.Crdb, error) {
	crdbCmd := exec.CommandContext(ctx, "cockroach",
		"start-single-node",
		"--store=type=mem",
		"--listen-addr=127.0.0.1:26257")
	if err := crdbCmd.Start(); err != nil {
		panic(fmt.Sprintf("error starting cockroachdb: %v", err))
	}
	return crdb.New(ctx, "postgres://postgres@127.0.0.1:26257?sslmode=disable")
}
