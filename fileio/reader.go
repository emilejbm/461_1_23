package fileio

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/19chonm/461_1_23/logger"
)

const url_ch_size = 100 // Size of the buffer for the URL channel

func IsValidURL(url_str string) bool {
	// Returns true if the domain or the url is github.com or npmjs.com, false otherwise
	u, e := url.Parse(url_str)
	if e != nil {
		logger.DebugMsg("fileio: ", e.Error())
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
	logger.InfoMsg("fileio: read file", path)

	file, e := os.Open(path)
	if e != nil {
		logger.InfoMsg("fileio: ", e.Error())
		fmt.Fprintf(os.Stderr, "fileio: Error reading file: %s\n", e.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // The default split function is ScanLines
		logger.InfoMsg("fileio: read entry:", scanner.Text(), fmt.Sprintf("%t", IsValidURL(scanner.Text())))
		if IsValidURL(scanner.Text()) {
			ch <- scanner.Text()
		} else {
			// Abort entire process if there is an invalid URL in the file
			logger.DebugMsg(fmt.Sprintf("Error processing file, invalid url: %s\n", scanner.Text()))
			fmt.Fprintf(os.Stderr, "fileio: Error processing file, invalid url: %s\n", scanner.Text())
			os.Exit(1)
		}
	}

	if scanner.Err() != nil { // not sure if correct
		logger.DebugMsg("fileio: ", scanner.Err().Error())
		fmt.Fprintf(os.Stderr, "fileio: Error reading file: %s\n", e.Error())
		os.Exit(1)
	}

	file.Close()
	close(ch)
}

func MakeUrlChannel() chan string {
	return make(chan string, url_ch_size)
}
