/*
Creation of build command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds executable for app",
	Long:  "Completes any compilation needed, builds executable with 'run' as name",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		app := "go"
		arg0 := "build"
		arg1 := "-o"
		arg2 := "metric_cli"

		exec_output := exec.Command(app, arg0, arg1, arg2)
		stdout, err := exec_output.Output()

		if err != nil {
			fmt.Println("CLI: ", err.Error())
		} else {
			fmt.Println("CLI: Build succesful", string(stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
