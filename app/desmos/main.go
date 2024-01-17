package main

import (
	"fmt"
	"os"

	"github.com/desmos-labs/desmos/v6/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/desmos-labs/desmos/v6/app/desmos/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
