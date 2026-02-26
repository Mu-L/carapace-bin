package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "keytool",
	Short: "Key and Certificate Management Tool",
	Long:  "https://github.com/openjdk/jdk/blob/master/src/java.base/share/man/keytool.md",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")

	embed.SubcommandsAsFlagsS(rootCmd,
		certreqCmd,
		changealiasCmd,
		deleteCmd,
		exportcertCmd,
		gencertCmd,
		genkeypairCmd,
		genseckeyCmd,
		importcertCmd,
		importkeystoreCmd,
		importpassCmd,
		keypasswdCmd,
		listCmd,
		printcertCmd,
		printcertreqCmd,
		printcrlCmd,
		showinfoCmd,
		storepasswdCmd,
		versionCmd,
	)
}
