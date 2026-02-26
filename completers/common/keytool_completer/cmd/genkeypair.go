package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genkeypairCmd = &cobra.Command{
	Use:     "genkeypair",
	Aliases: []string{"genkeypair"},
	Short:   "Generates a key pair",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genkeypairCmd).Standalone()

	genkeypairCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	genkeypairCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	genkeypairCmd.Flags().StringS("dname", "dname", "", "distinguished name")
	genkeypairCmd.Flags().StringS("ext", "ext", "", "X.509 extension")
	genkeypairCmd.Flags().StringS("groupname", "groupname", "", "Group name. For example, an Elliptic Curve name.")
	genkeypairCmd.Flags().StringS("keyalg", "keyalg", "", "key algorithm name")
	genkeypairCmd.Flags().StringS("keypass", "keypass", "", "key password")
	genkeypairCmd.Flags().StringS("keysize", "keysize", "", "key bit size")
	genkeypairCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	genkeypairCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	genkeypairCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	genkeypairCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	genkeypairCmd.Flags().StringS("providername", "providername", "", "provider name")
	genkeypairCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	genkeypairCmd.Flags().StringS("sigalg", "sigalg", "", "signature algorithm name")
	genkeypairCmd.Flags().StringS("signer", "signer", "", "signer alias")
	genkeypairCmd.Flags().StringS("signerkeypass", "signerkeypass", "", "signer key password")
	genkeypairCmd.Flags().StringS("startdate", "startdate", "", "certificate validity start date/time")
	genkeypairCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	genkeypairCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	genkeypairCmd.Flags().BoolS("v", "v", false, "verbose output")
	genkeypairCmd.Flags().StringS("validity", "validity", "", "validity number of days")
}
