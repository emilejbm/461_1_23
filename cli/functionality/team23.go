/*
This file contains extra functionality for CLI commands. Make sure to
capitalize first letter of functions to be used in external folders.
*/

package functionality

import (
	"fmt"
	"os"
)

func test(input string) int {
	return 0
}

func Read_url_file(filename string) int {

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("CLI: File reading error", err)
		os.Exit(1)
	}
	fmt.Println("CLI: Contents of file:", string(contents))

	return 0
}

func build() int {
	return 0
}
