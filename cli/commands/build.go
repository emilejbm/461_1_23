/*
Creation of build command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os/exec"

	"github.com/19chonm/461_1_23/logger"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds executable for app",
	Long:  "Completes any compilation needed, builds executable with 'run' as name",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		app := "go"
		arg := []string{"build", "-o", "metric_cli"}

		exec_output := exec.Command(app, arg...)
		stdout, err := exec_output.CombinedOutput()

		if err != nil {
			logger.DebugMsg("CLI: ", fmt.Sprint(err)+": "+string(stdout))
		} else {
			logger.InfoMsg("CLI: Build successful", string(stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
