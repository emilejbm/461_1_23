package main

import "github.com/19chonm/461_1_23/cli/commands"

func main() {
	commands.Execute()
}

// import (
// 	"fmt"
// 	"sync"
// )

// func doTask(url string) {
// 	fmt.Println("My job is", url)
// }

// func startWorkers(workers int, ch <-chan string) {
// 	var wg sync.WaitGroup

// 	// This starts workers number of goroutines that wait for something to do
// 	wg.Add(workers)
// 	for i := 0; i < workers; i++ {
// 		go func() {
// 			for {
// 				url, ok := <-ch
// 				// fmt.Println("taken: " + url)
// 				if !ok { // Channel has been closed
// 					// fmt.Println("channel closed")
// 					wg.Done()
// 					return
// 				}
// 				doTask(url)
// 			}
// 		}()
// 	}

// 	wg.Wait() // Wait for the threads to finish
// }

// func main() {
// 	url_ch := make(chan string, 5) // change buffer to 100
// 	go readFiles("urls_test.txt", url_ch)

// 	const workers = 5 // Total number of workers/goroutines to use
// 	startWorkers(workers, url_ch)
// }

// //https://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time
// //https://stackoverflow.com/questions/55203251/limiting-number-of-go-routines-running
