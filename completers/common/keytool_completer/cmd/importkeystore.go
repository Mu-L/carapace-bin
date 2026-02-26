package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importkeystoreCmd = &cobra.Command{
	Use:     "importkeystore",
	Aliases: []string{"importkeystore"},
	Short:   "Imports one or all entries from another keystore",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importkeystoreCmd).Standalone()

	importkeystoreCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	importkeystoreCmd.Flags().StringS("destalias", "destalias", "", "destination alias")
	importkeystoreCmd.Flags().StringS("destkeypass", "destkeypass", "", "destination key password")
	importkeystoreCmd.Flags().StringS("destkeystore", "destkeystore", "", "destination keystore name")
	importkeystoreCmd.Flags().BoolS("destprotected", "destprotected", false, "destination keystore password protected")
	importkeystoreCmd.Flags().StringS("destprovidername", "destprovidername", "", "destination keystore provider name")
	importkeystoreCmd.Flags().StringS("deststorepass", "deststorepass", "", "destination keystore password")
	importkeystoreCmd.Flags().StringS("deststoretype", "deststoretype", "", "destination keystore type")
	importkeystoreCmd.Flags().BoolS("noprompt", "noprompt", false, "do not prompt")
	importkeystoreCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	importkeystoreCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	importkeystoreCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	importkeystoreCmd.Flags().StringS("srcalias", "srcalias", "", "source alias")
	importkeystoreCmd.Flags().StringS("srckeypass", "srckeypass", "", "source key password")
	importkeystoreCmd.Flags().StringS("srckeystore", "srckeystore", "", "source keystore name")
	importkeystoreCmd.Flags().BoolS("srcprotected", "srcprotected", false, "source keystore password protected")
	importkeystoreCmd.Flags().StringS("srcprovidername", "srcprovidername", "", "source keystore provider name")
	importkeystoreCmd.Flags().StringS("srcstorepass", "srcstorepass", "", "source keystore password")
	importkeystoreCmd.Flags().StringS("srcstoretype", "srcstoretype", "", "source keystore type")
	importkeystoreCmd.Flags().BoolS("v", "v", false, "verbose output")
}
