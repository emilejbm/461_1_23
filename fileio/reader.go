package fileio

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const url_ch_size = 100 // Size of the buffer for the URL channel

func isValidURL(url_str string) bool {
	// Returns true if the domain or the url is github.com or npmjs.com, false otherwise
	u, e := url.Parse(url_str)
	if e != nil {
		fmt.Println("fileio: ", e.Error())
		return false
	}

	hostname := strings.TrimPrefix(u.Hostname(), "www.")

	if hostname == "github.com" || hostname == "npmjs.com" {
		return true
	}

	return false
}

func ReadFile(path string, ch chan<- string) {
	// Reads the file from path, parses and verifies the URL,
	// then sends valid URLs to channel ch
	fmt.Println("fileio: read file", path)

	file, e := os.Open(path)
	if e != nil {
		fmt.Println("fileio: ", e.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // The default split function is ScanLines
		fmt.Println("fileio: read entry:", scanner.Text(), isValidURL(scanner.Text()))
		if isValidURL(scanner.Text()) {
			ch <- scanner.Text()
		} else {
			// Abort entire process if there is an invalid URL in the file
			fmt.Fprintf(os.Stderr, "Error processing file, invalid url: %s\n", scanner.Text())
			os.Exit(1)
		}
	}

	if scanner.Err() != nil { // not sure if correct
		fmt.Println("fileio: ", scanner.Err())
	}

	file.Close()
	close(ch)
}

func MakeUrlChannel() chan string {
	return make(chan string, url_ch_size)
}
