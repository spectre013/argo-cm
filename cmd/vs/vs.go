package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.platform1.ninja/platform/furious-five/platform-orchestration/platform-deployment/argo-utils/virtualservice"
)


var vs = virtualservice.NewVirtualService()

var Command = &cobra.Command{
	Use:   "vs",
	Short: "Variable replacement on the virtualservice.yaml",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		vars, _ := cmd.Flags().GetString("vars")
		mvp, _ := cmd.Flags().GetBool("mvp")
		return vs.Process(vars, mvp)
	},
}

func init() {
	Command.Flags().StringP("vars", "v","", "Comma seperated list of env variables to replace")
	Command.Flags().BoolP("mvp", "m",false, "is MVP")
}