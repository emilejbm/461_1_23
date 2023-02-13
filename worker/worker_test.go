package worker

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
)

// test with bad url
func TestWorkerBadInput(t *testing.T) {
	worker_output_ch := make(chan fileio.WorkerOutput, 1)
	runTask("https://badurl.com/blabla/test", worker_output_ch)
	close(worker_output_ch)
	output := <-worker_output_ch
	if output.WorkerErr == nil {
		t.Errorf("rating channel should not have been updated")
	}
}

// test with incorrect owner for repo
func TestWorkerRatingFail(t *testing.T) {
	worker_output_ch := make(chan fileio.WorkerOutput, 1)

	incorrectUrl := "https://github.com/incorrectownername/react"
	runTask(incorrectUrl, worker_output_ch)
	close(worker_output_ch)

	for wo := range worker_output_ch {
		rating := wo.WorkerRating
		if rating.NetScore != 0 {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}

// tests for a package that should have net score of 0 because of license
func TestWorkerLicenseFail(t *testing.T) {
	worker_output_ch := make(chan fileio.WorkerOutput, 1)

	incorrectUrl := "https://github.com/expressjs/express"
	runTask(incorrectUrl, worker_output_ch)
	close(worker_output_ch)
	for wo := range worker_output_ch {
		rating := wo.WorkerRating
		if rating.NetScore != 0 && rating.License == 0 {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}
