package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// check how stackoverflow works again
// https://gosamples.dev/get-hostname-domain/
func isValidURL(url_str string) bool {
	url, err := url.Parse(url_str)
	if err != nil {
		fmt.Println(err.Error())
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")

	if hostname == "github.com" || hostname == "npmjs.com" {
		return true
	}

	return false
}

func readFiles(filePath string, ch chan<- string) {

	// var filepath string = "urls_test.txt"

	f, e := os.Open(filePath)
	if e != nil {
		panic(e.Error())
	}

	scanner := bufio.NewScanner(f) // default split function is ScanLines (https://pkg.go.dev/bufio#NewScanner)
	for scanner.Scan() {
		fmt.Println(scanner.Text(), isValidURL(scanner.Text()))
		if isValidURL(scanner.Text()) {
			ch <- scanner.Text()
		}
		// fmt.Println("add: ", scanner.Text())
	}

	if scanner.Err() != nil { // not sure if correct
		fmt.Printf("error: %s\n", scanner.Err())
	}

	f.Close()
	close(ch)
}

// Unit tests:
// Handle empty newline at end of file
// Handle lack of newline at end of file

// bash-4.2$ npm repo axios --browser false
// axios repo available at the following URL:
//   https://github.com/axios/axios

// bash-4.2$ npm repo connect --browser false
// connect repo available at the following URL:
//   https://github.com/senchalabs/connect

// https://stackoverflow.com/questions/34071621/query-npmjs-registry-via-api
