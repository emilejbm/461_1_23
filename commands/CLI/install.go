package team23

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use: "install",
	Short: "install dependencies to use CLI and app",
	Long: "installs all dependencies required to use CLI and app",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string){
		fmt.Println("install command recognized")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	fmt.Println("install command init")
}