package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printcertreqCmd = &cobra.Command{
	Use:     "printcertreq",
	Aliases: []string{"printcertreq"},
	Short:   "Prints the content of a certificate request",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printcertreqCmd).Standalone()

	printcertreqCmd.Flags().StringS("file", "file", "", "input file name")
	printcertreqCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(printcertreqCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
