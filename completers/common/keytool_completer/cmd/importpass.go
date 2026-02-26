package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importpassCmd = &cobra.Command{
	Use:     "importpass",
	Aliases: []string{"importpass"},
	Short:   "Imports a password",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importpassCmd).Standalone()

	importpassCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	importpassCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	importpassCmd.Flags().StringS("keyalg", "keyalg", "", "key algorithm name")
	importpassCmd.Flags().StringS("keypass", "keypass", "", "key password")
	importpassCmd.Flags().StringS("keysize", "keysize", "", "key bit size")
	importpassCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	importpassCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	importpassCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	importpassCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	importpassCmd.Flags().StringS("providername", "providername", "", "provider name")
	importpassCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	importpassCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	importpassCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	importpassCmd.Flags().BoolS("v", "v", false, "verbose output")
}
