package main

import (
	"os"

	"github.com/desmos-labs/desmos/app/desmosd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
