package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certreqCmd = &cobra.Command{
	Use:     "certreq",
	Aliases: []string{"certreq"},
	Short:   "Generates a certificate request",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certreqCmd).Standalone()

	certreqCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	certreqCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	certreqCmd.Flags().StringS("dname", "dname", "", "distinguished name")
	certreqCmd.Flags().StringS("ext", "ext", "", "X.509 extension")
	certreqCmd.Flags().StringS("file", "file", "", "output file name")
	certreqCmd.Flags().StringS("keypass", "keypass", "", "key password")
	certreqCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	certreqCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	certreqCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	certreqCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	certreqCmd.Flags().StringS("providername", "providername", "", "provider name")
	certreqCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	certreqCmd.Flags().StringS("sigalg", "sigalg", "", "signature algorithm name")
	certreqCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	certreqCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	certreqCmd.Flags().BoolS("v", "v", false, "verbose output")

	carapace.Gen(certreqCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
