package worker

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
)

var url_ch_size = 100
var wo_ch_size = 5

func TestManagerBadInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	worker_output_ch := make(chan fileio.WorkerOutput, wo_ch_size)
	worker_outputs := []fileio.WorkerOutput{}
	url_ch <- "https://badurl.com/blabla/test"
	url_ch <- "incorrecturl!!@#$**F(S)"

	go StartWorkers(url_ch, worker_output_ch)
	close(url_ch)

	for {
		r, ok := <-worker_output_ch
		if !ok { // Channel has been closed
			break
		}
		worker_outputs = append(worker_outputs, r)
	}

	// bad url should make a rating of default values
	for _, wo := range worker_outputs {
		err := wo.WorkerErr
		if err == nil {
			t.Errorf("worker should have errored for bad url")
		}
	}
}
