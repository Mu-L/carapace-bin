package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_apiserverKubeletClientCmd = &cobra.Command{
	Use:   "apiserver-kubelet-client",
	Short: "Generate the certificate for the API server to connect to kubelet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_apiserverKubeletClientCmd).Standalone()

	init_phase_certs_apiserverKubeletClientCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_certs_apiserverKubeletClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_apiserverKubeletClientCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_certs_apiserverKubeletClientCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_apiserverKubeletClientCmd)

	carapace.Gen(init_phase_certs_apiserverKubeletClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
