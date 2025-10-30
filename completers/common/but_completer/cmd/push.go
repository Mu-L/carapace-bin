package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push BRANCH_ID",
	Short: "Push a branch/stack to remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolP("help", "h", false, "Print help")
	pushCmd.Flags().BoolP("run-hooks", "r", false, "Run pre-push hooks")
	pushCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	pushCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
