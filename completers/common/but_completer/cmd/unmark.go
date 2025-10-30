package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unmarkCmd = &cobra.Command{
	Use:   "unmark",
	Short: "Removes all marks from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmarkCmd).Standalone()

	unmarkCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(unmarkCmd)
}
