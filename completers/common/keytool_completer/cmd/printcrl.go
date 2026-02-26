package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printcrlCmd = &cobra.Command{
	Use:     "printcrl",
	Aliases: []string{"printcrl"},
	Short:   "Prints the content of a CRL file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printcrlCmd).Standalone()

	printcrlCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	printcrlCmd.Flags().StringS("file", "file", "", "input file name")
	printcrlCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	printcrlCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	printcrlCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	printcrlCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	printcrlCmd.Flags().StringS("providername", "providername", "", "provider name")
	printcrlCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	printcrlCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	printcrlCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	printcrlCmd.Flags().BoolS("trustcacerts", "trustcacerts", false, "trust certificates from cacerts")
	printcrlCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(printcrlCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
