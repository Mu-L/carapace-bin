package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importcertCmd = &cobra.Command{
	Use:     "importcert",
	Aliases: []string{"importcert"},
	Short:   "Imports a certificate or a certificate chain",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importcertCmd).Standalone()

	importcertCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	importcertCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	importcertCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	importcertCmd.Flags().StringS("file", "file", "", "input file name")
	importcertCmd.Flags().StringS("keypass", "keypass", "", "key password")
	importcertCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	importcertCmd.Flags().BoolS("noprompt", "noprompt", false, "do not prompt")
	importcertCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	importcertCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	importcertCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	importcertCmd.Flags().StringS("providername", "providername", "", "provider name")
	importcertCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	importcertCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	importcertCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	importcertCmd.Flags().BoolS("trustcacerts", "trustcacerts", false, "trust certificates from cacerts")
	importcertCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(importcertCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
