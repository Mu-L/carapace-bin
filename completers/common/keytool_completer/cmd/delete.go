package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"delete"},
	Short:   "Deletes an entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	deleteCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	deleteCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	deleteCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	deleteCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	deleteCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	deleteCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	deleteCmd.Flags().StringS("providername", "providername", "", "provider name")
	deleteCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	deleteCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	deleteCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	deleteCmd.Flags().BoolS("v", "v", false, "verbose output")
}
