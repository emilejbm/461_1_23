package fileio

import (
	"encoding/json"
	"fmt"
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

const rating_ch_size = 100 // Size of the buffer for the URL channel

type Rating struct {
	Url            string  `json:"URL"`
	NetScore       float64 `json:"NET_SCORE"`
	Rampup         float64 `json:"RAMP_UP_SCORE"`
	Correctness    float64 `json:"CORRECTNESS_SCORE"`
	Busfactor      float64 `json:"BUS_FACTOR_SCORE"`
	Responsiveness float64 `json:"RESPONSIVE_MAINTAINER_SCORE"`
	License        float64 `json:"LICENSE_SCORE"`
}

func MakeRatingsChannel() chan Rating {
	return make(chan Rating, rating_ch_size)
}

func Sort_modules(ch chan Rating) []Rating {
	// Create a slice to hold the values from the channel
	sorted_ratings := []Rating{}

	// Read in ratings from channel
	for {
		r, ok := <-ch
		if !ok { // Channel has been closed
			break
		}
		sorted_ratings = append(sorted_ratings, r)
	}

	// Sort the slice
	sort.Slice(sorted_ratings, func(p, q int) bool {
		return sorted_ratings[p].NetScore > sorted_ratings[q].NetScore
	})

	return sorted_ratings
}

func Make_json_string(r Rating) string {
	// Convert the Rating struct into a json string
	jsonString, err := json.Marshal(r)
	if err != nil {
		logger.DebugMsg(fmt.Sprintf("for: %+v\n", r), "fileio: make_json_string fail")
		fmt.Printf("for: %+v\n", r)
		panic("fileio: Make_json_string fail")
	}

	return string(jsonString)
}

func Print_sorted_output(ratings []Rating) {
	for i := range ratings {
		logger.InfoMsg(Make_json_string(ratings[i]))
	}
}
