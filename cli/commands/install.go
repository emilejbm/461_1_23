/*
Creation of install command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os/exec"

	"github.com/19chonm/461_1_23/logger"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dependencies to use app",
	Long:  "As of now, only installs cobra package which is required to use our Go app",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// Install Logrus
		app := "go"
		arg := []string{"get", "github.com/sirupsen/logrus"}
		exec_output := exec.Command(app, arg...)
		stdout, err := exec_output.Output()

		if err != nil {
			logger.DebugMsg("CLI: ", err.Error())
			return
		} else {
			logger.InfoMsg("CLI: Installed Logrus", string(stdout))
		}

		// Install cobra
		app = "go"
		arg = []string{"get", "-u", "github.com/spf13/cobra@latest"}

		exec_output = exec.Command(app, arg...)
		stdout, err = exec_output.Output()

		if err != nil {
			logger.DebugMsg("CLI: ", err.Error())
			return
		} else {
			logger.InfoMsg("CLI: Installed Cobra", string(stdout))
		}

		// Make Python venv
		logger.InfoMsg("CLI: Make python venv")
		app = "python3"
		arg = []string{"-m", "venv", "venv"}

		exec_output = exec.Command(app, arg...)
		stdout, err = exec_output.CombinedOutput()

		if err != nil {
			logger.DebugMsg("CLI: ", fmt.Sprint(err)+": "+string(stdout))
			return
		} else {
			logger.InfoMsg("CLI: Python venv complete", string(stdout))
		}

		// Install GitPython
		logger.InfoMsg("CLI: Installing GitPython")
		app = "./venv/bin/pip"
		arg = []string{"install", "gitpython"}

		exec_output = exec.Command(app, arg...)
		stdout, err = exec_output.CombinedOutput()

		if err != nil {
			logger.DebugMsg("CLI: ", fmt.Sprint(err)+": "+string(stdout))
			return
		} else {
			logger.InfoMsg("CLI: Installed GitPython", string(stdout))
		}

		logger.InfoMsg("CLI: Installation successful")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
