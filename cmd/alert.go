package cmd

import (
	"github.com/anynines/a9s-aha-cli/pkg/aha"

	"github.com/spf13/cobra"
)

func newCmdAlert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alert",
		Short: "Alert on command line if any alerts present",
		Long:  `Alert on command line if any alerts present.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			notifyFlag, _ := cmd.Flags().GetBool("notify")
			verboseFlag, _ := cmd.Flags().GetBool("verbose")

			return aha.Stale(notifyFlag, verboseFlag)
		},
	}

	cmd.Flags().BoolP("notify", "n", false, "Notify on Slack.")
	cmd.Flags().BoolP("verbose", "v", false, "Provide additional debug details.")
	return cmd
}
