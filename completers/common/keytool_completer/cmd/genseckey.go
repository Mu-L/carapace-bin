package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genseckeyCmd = &cobra.Command{
	Use:     "genseckey",
	Aliases: []string{"genseckey"},
	Short:   "Generates a secret key",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genseckeyCmd).Standalone()

	genseckeyCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	genseckeyCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	genseckeyCmd.Flags().StringS("keyalg", "keyalg", "", "key algorithm name")
	genseckeyCmd.Flags().StringS("keypass", "keypass", "", "key password")
	genseckeyCmd.Flags().StringS("keysize", "keysize", "", "key bit size")
	genseckeyCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	genseckeyCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	genseckeyCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	genseckeyCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	genseckeyCmd.Flags().StringS("providername", "providername", "", "provider name")
	genseckeyCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	genseckeyCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	genseckeyCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	genseckeyCmd.Flags().BoolS("v", "v", false, "verbose output")
}
