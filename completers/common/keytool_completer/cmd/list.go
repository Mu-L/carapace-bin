package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "Lists entries in a keystore",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	listCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	listCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	listCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	listCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	listCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	listCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	listCmd.Flags().StringS("providername", "providername", "", "provider name")
	listCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	listCmd.Flags().BoolS("rfc", "rfc", false, "output in RFC style")
	listCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	listCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	listCmd.Flags().BoolS("v", "v", false, "verbose output")
}
