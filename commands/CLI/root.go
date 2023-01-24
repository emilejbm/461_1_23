package team23

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "team23",
	Short: "team23 - root command for app",
	Long: "team23 is the root command to navigate through Team 23's CLI app",
	Run: func(cmd *cobra.Command, args[]string){

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error using CLI '%s'", err)
		os.Exit(1)
	}
}