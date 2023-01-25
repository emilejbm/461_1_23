package team23

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/emilejbm/461_1_23/functionality/CLI"
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
		functionality.READ_url_file(os.Args)
	},
}

func init() {
	rootCmd.AddCommand(urlfileCmd)
	//fmt.Println("url_file command init")
}