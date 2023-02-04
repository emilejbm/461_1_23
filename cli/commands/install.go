/*
Creation of install command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dependencies to use app",
	Long:  "As of now, only installs cobra package which is required to use our Go app",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// Install cobra
		app := "go"
		arg := []string{"get", "-u", "github.com/spf13/cobra@latest"}

		exec_output := exec.Command(app, arg...)
		stdout, err := exec_output.Output()

		if err != nil {
			fmt.Println("CLI: ", err.Error())
			return
		}

		// Make Python venv
		fmt.Println("CLI: Make python venv")
		app = "python3"
		arg = []string{"-m", "venv", "venv"}

		exec_output = exec.Command(app, arg...)
		stdout, err = exec_output.CombinedOutput()

		if err != nil {
			fmt.Println("CLI: ", fmt.Sprint(err)+": "+string(stdout))
			return
		} else {
			fmt.Println("CLI: Python venv complete", string(stdout))
		}

		// Install GitPython
		fmt.Println("CLI: Installing GitPython")
		app = "./venv/bin/pip"
		arg = []string{"install", "gitpython"}

		exec_output = exec.Command(app, arg...)
		stdout, err = exec_output.CombinedOutput()

		if err != nil {
			fmt.Println("CLI: ", fmt.Sprint(err)+": "+string(stdout))
			return
		} else {
			fmt.Println("CLI: Installed GitPython", string(stdout))
		}

		fmt.Println("CLI: Installation succesful")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
