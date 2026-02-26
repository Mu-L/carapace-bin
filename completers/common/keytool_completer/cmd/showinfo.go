package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showinfoCmd = &cobra.Command{
	Use:     "showinfo",
	Aliases: []string{"showinfo"},
	Short:   "Displays security-related information",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showinfoCmd).Standalone()

	showinfoCmd.Flags().BoolS("tls", "tls", false, "Displays TLS configuration information")
	showinfoCmd.Flags().BoolS("v", "v", false, "verbose output")
}
