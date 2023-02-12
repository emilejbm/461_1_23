package worker

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
)

// success test
func TestWorkerPositiveScore(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)
	runTask("https://github.com/nullivex/nodist", rating_ch)
	close(rating_ch)
	for rating := range rating_ch {
		if rating.NetScore == 0 {
			t.Errorf("rating should be more than 0")
		}
	}
}

// test with bad url
func TestWorkerBadInput(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)
	runTask("https://badurl.com/blabla/test", rating_ch)
	close(rating_ch)
	if len(rating_ch) != 0 {
		t.Errorf("rating channel should not have been updated")
	}
}

// test with incorrect owner for repo
func TestWorkerRatingFail(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)

	incorrectUrl := "https://github.com/incorrectownername/react"
	runTask(incorrectUrl, rating_ch)
	close(rating_ch)

	for rating := range rating_ch {
		if rating.NetScore != 0 {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}

// tests for a package that should have net score of 0 because of license
func TestWorkerLicenseFail(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)

	incorrectUrl := "https://github.com/expressjs/express"
	runTask(incorrectUrl, rating_ch)
	close(rating_ch)
	for rating := range rating_ch {
		if rating.NetScore != 0 && rating.License == 0 {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}
