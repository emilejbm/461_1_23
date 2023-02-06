package worker

import (
	"fmt"
	"math/rand"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/metrics"
)

func runTask(url string, ratingch chan<- fileio.Rating) {
	fmt.Println("My job is", url)
	rampupscore := metrics.ScanRepo(url)
	r := fileio.Rating{NetScore: rand.Float64(), Rampup: rampupscore, Url: url} // Placeholder until real data implemented
	ratingch <- r
}
