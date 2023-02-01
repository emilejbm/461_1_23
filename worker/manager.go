package worker

import (
	"fmt"
	"sync"
)

const numworkers = 5 // Total number of workers/goroutines to use

func StartWorkers(ch <-chan string) {
	var wg sync.WaitGroup

	wg.Add(numworkers) // Keep track of the number of goroutines being created
	for i := 0; i < numworkers; i++ {
		// Start each worker
		go func() {
			fmt.Println("worker: Start worker")
			for {
				url, ok := <-ch
				if !ok { // Channel has been closed
					fmt.Println("worker: Close worker")
					wg.Done()
					return
				}
				runTask(url)
			}
		}()
	}

	wg.Wait() // Wait for the threads to finish
}
