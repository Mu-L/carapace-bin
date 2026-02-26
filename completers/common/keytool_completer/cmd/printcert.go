package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printcertCmd = &cobra.Command{
	Use:     "printcert",
	Aliases: []string{"printcert"},
	Short:   "Prints the content of a certificate",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printcertCmd).Standalone()

	printcertCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	printcertCmd.Flags().StringS("file", "file", "", "input file name")
	printcertCmd.Flags().StringS("jarfile", "jarfile", "", "signed jar file")
	printcertCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	printcertCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	printcertCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	printcertCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	printcertCmd.Flags().StringS("providername", "providername", "", "provider name")
	printcertCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	printcertCmd.Flags().BoolS("rfc", "rfc", false, "output in RFC style")
	printcertCmd.Flags().StringS("sslserver", "sslserver", "", "SSL server host and port")
	printcertCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	printcertCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	printcertCmd.Flags().BoolS("trustcacerts", "trustcacerts", false, "trust certificates from cacerts")
	printcertCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(printcertCmd).FlagCompletion(carapace.ActionMap{
		"file":    carapace.ActionFiles(),
		"jarfile": carapace.ActionFiles(),
	})
}
