package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	vs "gitlab.platform1.ninja/platform/furious-five/platform-orchestration/platform-deployment/argo-utils/cmd/vs"
)

var cfgFile string



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kweng",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Hello")
	//
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(vs.Command)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
