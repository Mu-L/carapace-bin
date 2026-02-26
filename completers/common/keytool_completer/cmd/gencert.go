package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gencertCmd = &cobra.Command{
	Use:     "gencert",
	Aliases: []string{"gencert"},
	Short:   "Generates certificate from a certificate request",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gencertCmd).Standalone()

	gencertCmd.Flags().StringS("addprovider", "addprovider", "", "add security provider by name (e.g. SunPKCS11)")
	gencertCmd.Flags().StringS("alias", "alias", "", "alias name of the entry to process")
	gencertCmd.Flags().StringS("dname", "dname", "", "distinguished name")
	gencertCmd.Flags().StringS("ext", "ext", "", "X.509 extension")
	gencertCmd.Flags().StringS("infile", "infile", "", "input file name")
	gencertCmd.Flags().StringS("keypass", "keypass", "", "key password")
	gencertCmd.Flags().StringS("keystore", "keystore", "", "keystore name")
	gencertCmd.Flags().StringS("outfile", "outfile", "", "output file name")
	gencertCmd.Flags().BoolS("protected", "protected", false, "password through protected mechanism")
	gencertCmd.Flags().StringArrayS("providerarg", "providerarg", nil, "configure argument for -addprovider/-providerclass")
	gencertCmd.Flags().StringS("providerclass", "providerclass", "", "add security provider by fully-qualified class name")
	gencertCmd.Flags().StringS("providername", "providername", "", "provider name")
	gencertCmd.Flags().StringS("providerpath", "providerpath", "", "provider classpath")
	gencertCmd.Flags().BoolS("rfc", "rfc", false, "output in RFC style")
	gencertCmd.Flags().StringS("sigalg", "sigalg", "", "signature algorithm name")
	gencertCmd.Flags().StringS("startdate", "startdate", "", "certificate validity start date/time")
	gencertCmd.Flags().StringS("storepass", "storepass", "", "keystore password")
	gencertCmd.Flags().StringS("storetype", "storetype", "", "keystore type")
	gencertCmd.Flags().BoolS("v", "v", false, "verbose output")
	gencertCmd.Flags().StringS("validity", "validity", "", "validity number of days")

	carapace.Gen(gencertCmd).FlagCompletion(carapace.ActionMap{
		"infile":  carapace.ActionFiles(),
		"outfile": carapace.ActionFiles(),
	})
}
