package fileio

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/19chonm/461_1_23/logger"
)

// 					--- NDJSON ---
// Each JSON text MUST conform to the [RFC8259]
// standard and MUST be written to the stream followed
// by the newline character \n (0x0A). The newline character
// MAY be preceded by a carriage return \r (0x0D). The JSON
// texts MUST NOT contain newlines or carriage returns.
//
// All serialized data MUST use the UTF8 encoding.
// https://github.com/ndjson/ndjson-spec

// ND JSON follows RFC 8259
// Go's JSON library follows RFC 7159
// - However, the only major change between the two is that 8259 supports UTF8, which Go does by default
// - Assuming this different is negligible, and that using Go's "json" library is okay

const worker_output_ch_size = 100 // Size of the buffer for the Worker's output

type Rating struct {
	Url            string  `json:"URL"`
	NetScore       float64 `json:"NET_SCORE"`
	Rampup         float64 `json:"RAMP_UP_SCORE"`
	Correctness    float64 `json:"CORRECTNESS_SCORE"`
	Busfactor      float64 `json:"BUS_FACTOR_SCORE"`
	Responsiveness float64 `json:"RESPONSIVE_MAINTAINER_SCORE"`
	License        float64 `json:"LICENSE_SCORE"`
}

type WorkerOutput struct {
	WorkerRating Rating
	WorkerErr    error
}

func MakeWorkerOutputChannel() chan WorkerOutput {
	return make(chan WorkerOutput, worker_output_ch_size)
}

func ReadWorkerResults(ch chan WorkerOutput) ([]Rating, []error) {
	// Create a slice to hold the values from the channel
	sorted_ratings := []Rating{}
	errors := []error{}

	// Read in ratings from channel
	for {
		wo, ok := <-ch
		if !ok { // Channel has been closed
			break
		}

		if wo.WorkerErr != nil { // If error, record the error
			errors = append(errors, wo.WorkerErr)
		} else { // Else, record the result
			sorted_ratings = append(sorted_ratings, wo.WorkerRating)
		}
	}

	// Sort the slice
	sort.Slice(sorted_ratings, func(p, q int) bool {
		return sorted_ratings[p].NetScore > sorted_ratings[q].NetScore
	})

	return sorted_ratings, errors
}

func Make_json_string(r Rating) string {
	// Convert the Rating struct into a json string
	jsonString, err := json.Marshal(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fileio: Make_json_string fail for: %+v\n", r)
		os.Exit(1)
	}

	return string(jsonString)
}

func Print_sorted_output(ratings []Rating) {
	for i := range ratings {
		logger.InfoMsg(Make_json_string(ratings[i]))
	}
}

func PrintErrors(errs []error) {
	fmt.Fprintln(os.Stderr, "Errors in the worker stage: ")
	for _, e := range errs {
		fmt.Fprintln(os.Stderr, e)
	}
}
