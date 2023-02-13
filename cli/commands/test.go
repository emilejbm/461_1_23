/*
Creation of test command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/19chonm/461_1_23/logger"

	"github.com/spf13/cobra"
	//"github.com/19chonm/461_1_23/api"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "runs test suite",
	Long: `Runs test suite located in -- exits 0 if everything is working.
	The minimum requirement for this test suite is that it contain at 
	least 20 distinct test cases and achieve at least 80'%' code coverage 
	as measured by line coverage. The output from this invocation should be 
	a line written to stdout of the form: “X/Y test cases passed. Z% line 
	coverage achieved.” 
	Should exit 0 on success.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//packagesToTest := []string{"api/", "worker/"}

		app := "go"
		testArgs := []string{"test", "-v", "-coverpkg=./...", "-coverprofile=profile.cov", "./..."}
		//testArgs := []string{"test", "./...", "-cover", "-v"}

		exec_output := exec.Command(app, testArgs...)
		stdout, err := exec_output.CombinedOutput()

		// distinction between "PASS" and "--- PASS" because if all tests cases pass,
		// "PASS" is printed out again
		output := string(stdout)
		testsPassed := strings.Count(output, "--- PASS")
		testsRan := strings.Count(output, "=== RUN")

		re := regexp.MustCompile(`(?P<coverage>coverage:\s)(?P<numbers>\d+\.\d)(?P<ofstatement>%\sof\sstatements in ./...)`)
		result := make(map[string]string)
		match := re.FindStringSubmatch(output)
		//fmt.Println(match)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		coverage := result["numbers"] + "%"

		if err != nil || testsPassed > testsRan {
			fmt.Printf("%d/%d test cases passed. %s line coverage achieved\n", testsPassed, testsRan, coverage)
			logger.DebugMsg("CLI: ", err.Error())
			os.Exit(1)
		}

		fmt.Printf("%d/%d test cases passed. %s line coverage achieved\n", testsPassed, testsRan, coverage)
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
