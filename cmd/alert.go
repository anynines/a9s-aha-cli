package cmd

import (
	"github.com/anynines/a9s-aha-cli/pkg/alert"

	"github.com/spf13/cobra"
)

func newCmdAlert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alert",
		Short: "Alert in Slack channel if any alerts present",
		Long:  `Alert in Slack channel if any alerts present.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return alert.Stale()
		},
	}

	return cmd
}
