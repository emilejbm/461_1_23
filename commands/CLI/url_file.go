package team23

import (
	"fmt"
	//"os"
	"github.com/spf13/cobra"
)

var urlfileCmd = &cobra.Command{
	Use: "URL_FILE",
	Short: "assign scores for packages found in URL_FILE",
	Long: `scans through the urls found within file passed as an argument
	and builds a score for the each package based on ramp-up time, correctness,
	bus factor, responsiveness from maintainer, and license compatibility. Will
	 produce a net score and output in NDJSON format to stdout.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args[]string){
		fmt.Println("url_file command recognized")
	},
}

func init() {
	rootCmd.AddCommand(urlfileCmd)
	fmt.Println("url_file command init")
}