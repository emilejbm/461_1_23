package team23

import (
	"fmt"
	//"github.com/emilejbm/461_1_23/functionality/CLI"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use: "test",
	Short: "runs test suite",
	Long: `runs test suite located in -- exits 0 if everything is working.
	The minimum requirement for this test suite is that it contain at least 20 distinct test cases and achieve at least 80'%' code coverage as measured by line coverage.
	The output from this invocation should be a line written to stdout of the form: “X/Y test cases passed. Z% line coverage achieved.” 
	Should exit 0 on success.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string){
		fmt.Println("test command recognized")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	fmt.Println("test command init")
}