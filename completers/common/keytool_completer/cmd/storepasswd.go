package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storepasswdCmd = &cobra.Command{
	Use:     "storepasswd",
	Aliases: []string{"storepasswd"},
	Short:   "Changes the store password of a keystore",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storepasswdCmd).Standalone()

	storepasswdCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	storepasswdCmd.Flags().BoolS("cacerts", "cacerts", false, "access the cacerts keystore")
	storepasswdCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	storepasswdCmd.Flags().StringS("new", "new", "", "new password")
	storepasswdCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	storepasswdCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	storepasswdCmd.Flags().StringS("providername", "providername", "", "provider name")
	storepasswdCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	storepasswdCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	storepasswdCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	storepasswdCmd.Flags().BoolS("v", "v", false, "verbose output")
}
