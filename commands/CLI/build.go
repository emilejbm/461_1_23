package team23

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use: "build",
	Short: "builds CLI app",
	Long: "completes any compilation needed",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string) {
		app := "go"
		arg0 := "build"
		arg1 := "main.go"

		exec_output := exec.Command(app, arg0, arg1)
		stdout, err := exec_output.Output()
		
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Build succesful", string(stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	//fmt.Println("build command init")
}