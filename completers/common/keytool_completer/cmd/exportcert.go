package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exportcertCmd = &cobra.Command{
	Use:     "exportcert",
	Aliases: []string{"exportcert"},
	Short:   "Exports certificate",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportcertCmd).Standalone()

	exportcertCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	exportcertCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	exportcertCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	exportcertCmd.Flags().StringS("file", "file", "", "output file name")
	exportcertCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	exportcertCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	exportcertCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	exportcertCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	exportcertCmd.Flags().StringS("providername", "providername", "", "provider name")
	exportcertCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	exportcertCmd.Flags().BoolS("rfc", "rfc", false, "output in RFC style")
	exportcertCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	exportcertCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	exportcertCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(exportcertCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
