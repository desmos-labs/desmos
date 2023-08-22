package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/desmos-labs/desmos/v6/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/desmos-labs/desmos/v6/app/desmos/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
