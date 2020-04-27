package main

import (
	"os"

	"github.com/anynines/a9s-aha-cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
