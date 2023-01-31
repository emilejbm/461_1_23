/*
Creation of install command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dependencies to use app",
	Long:  "As of now, only installs cobra package which is required to use our Go app",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		app := "go"
		arg0 := "get"
		arg1 := "-u"
		arg2 := "github.com/spf13/cobra@latest"

		exec_output := exec.Command(app, arg0, arg1, arg2)
		stdout, err := exec_output.Output()

		if err != nil {
			fmt.Println("CLI: ", err.Error())
		} else {
			fmt.Println("CLI: Installation succesful", string(stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
