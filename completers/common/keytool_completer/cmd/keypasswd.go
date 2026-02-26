package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keypasswdCmd = &cobra.Command{
	Use:     "keypasswd",
	Aliases: []string{"keypasswd"},
	Short:   "Changes the key password of an entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypasswdCmd).Standalone()

	keypasswdCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	keypasswdCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	keypasswdCmd.Flags().StringS("keypass", "keypass", "", "key password")
	keypasswdCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	keypasswdCmd.Flags().StringS("new", "new", "", "new password")
	keypasswdCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	keypasswdCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	keypasswdCmd.Flags().StringS("providername", "providername", "", "provider name")
	keypasswdCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	keypasswdCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	keypasswdCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	keypasswdCmd.Flags().BoolS("v", "v", false, "verbose output")
}
