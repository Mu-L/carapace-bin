package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changealiasCmd = &cobra.Command{
	Use:     "changealias",
	Aliases: []string{"changealias"},
	Short:   "Changes an entry's alias",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changealiasCmd).Standalone()

	changealiasCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	changealiasCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	changealiasCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	changealiasCmd.Flags().StringS("destalias", "destalias", "", "destination alias")
	changealiasCmd.Flags().StringS("keypass", "keypass", "", "key password")
	changealiasCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	changealiasCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	changealiasCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	changealiasCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	changealiasCmd.Flags().StringS("providername", "providername", "", "provider name")
	changealiasCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	changealiasCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	changealiasCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	changealiasCmd.Flags().BoolS("v", "v", false, "verbose output")
}
