package team23

import (
	"fmt"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use: "build",
	Short: "builds CLI app",
	Long: "completes any compilation needed",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string){
		fmt.Println("build command recognized")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	fmt.Println("build command init")
}