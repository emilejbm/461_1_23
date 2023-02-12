package worker

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
)

var url_ch_size = 100
var rating_ch_size = 5

func TestManagerGoodInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	ratings := []fileio.Rating{}
	url_ch <- "https://github.com/axios/axios"
	url_ch <- "https://github.com/nullivex/nodist"
	url_ch <- "https://github.com/cloudinary/cloudinary_npm"

	go StartWorkers(url_ch, rating_ch)
	close(url_ch)

	for {
		r, ok := <-rating_ch
		if !ok { // Channel has been closed
			break
		}
		ratings = append(ratings, r)
	}

	// comparison to default values of rating type to check
	// if they were untouched
	for _, rating := range ratings {
		if rating.Busfactor == 0 && rating.Correctness == 0 && rating.License == 0 && rating.NetScore == 0 && rating.Rampup == 0 && rating.Responsiveness == 0 {
			t.Errorf("ratings were not created correctly")
		}
	}
}

func TestManagerBadInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	ratings := []fileio.Rating{}
	url_ch <- "https://badurl.com/blabla/test"
	url_ch <- "incorrecturl!!@#$**F(S)"

	go StartWorkers(url_ch, rating_ch)
	close(url_ch)

	for {
		r, ok := <-rating_ch
		if !ok {
			break
		}
		ratings = append(ratings, r)
	}

	// bad url should make a rating of default values
	for _, rating := range ratings {
		if rating.Busfactor != 0 && rating.Correctness != 0 && rating.License != 0 && rating.NetScore != 0 && rating.Rampup != 0 && rating.Responsiveness != 0 {
			t.Errorf("ratings should not have been created")
		}
	}
}
