package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Creates or removes a rule for auto-assigning or auto-comitting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markCmd).Standalone()

	markCmd.Flags().BoolP("delete", "d", false, "Deletes a mark")
	markCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(markCmd)
}
