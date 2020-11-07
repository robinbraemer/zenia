package testing

import (
	"context"
	"fmt"
	"os/exec"
)

// use "defer Stop()" to stop the DB at end of test
var Ctx, Stop = context.WithCancel(context.Background())

func init() {
	crdbCmd := exec.CommandContext(Ctx, "cockroachdb",
		"start-single-node",
		"--store=type=mem")
	if err := crdbCmd.Start(); err != nil {
		panic(fmt.Sprintf("error starting cockroachdb: %v", err))
	}
}
