/*
Root command is created as team23. Holds no functionality but all other
commands are built on top of this command. Creation of new commands
requires an init function per command with rootCmd.AddCommand(<newCmd>)
*/

package commands

import (
	"fmt"
	"github.com/19chonm/461_1_23/cli/functionality"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "team23",
	Short: "team23 - root command for app",
	Long:  "team23 is the root command to navigate through Team 23's CLI",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// test URL_FILE with: /Users/emile/461_1_23/test/urls_file.txt

// First function to be ran on main. Will check if second argument is either
// an absolute filepath, one of the recognized commands or neither. If neither,
// program will throw error. If argument is an absolute filepath, a direct call
// to functions are executed. No cobra command is created because name varies.
func Execute() {

	if len(os.Args) != 2 {
		fmt.Println(`CLI: Please use one of the recognized commands: 'build', 
		'install', 'test', or 'URL_FILE' where URL_FILE is an absolute path 
		to a file`)
	} else if filepath.IsAbs(os.Args[1]) {
		functionality.Read_url_file(os.Args[1])
	} else if os.Args[1] == "build" || os.Args[1] == "install" ||
		os.Args[1] == "test" {

		if err := rootCmd.Execute(); err != nil {
			fmt.Println("CLI: Error using CLI '%s'", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("CLI: Not a recognized command\n")
		os.Exit(1)
	}

	os.Exit(0)
}
