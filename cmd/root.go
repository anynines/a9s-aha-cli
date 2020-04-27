package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the root Cobra command
var RootCmd = &cobra.Command{
	Use:   "a9s-aha-cli",
	Short: "a9s-aha-cli is an CLI to aha.io with features interesting to a9s organisation",
	Long:  `a9s-aha-cli is an CLI to aha.io with features interesting to a9s organisation.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(newCmdAlert())
}
