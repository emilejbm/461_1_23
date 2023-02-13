package worker

import (
	"sync"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/logger"
)

const numworkers = 5 // Total number of workers/goroutines to use

func StartWorkers(urlch <-chan string, worker_output_ch chan<- fileio.WorkerOutput) {
	var wg sync.WaitGroup

	wg.Add(numworkers) // Keep track of the number of goroutines being created
	for i := 0; i < numworkers; i++ {
		// Start each worker
		go func() {
			logger.InfoMsg("worker: Start worker")
			for {
				url, ok := <-urlch
				if !ok { // Channel has been closed
					logger.InfoMsg(("worker: Close worker"))
					wg.Done()
					return
				}
				runTask(url, worker_output_ch)
			}
		}()
	}

	wg.Wait() // Wait for the threads to finish
	close(worker_output_ch)
}
