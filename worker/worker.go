package worker

import (
	"fmt"
	"math/rand"

	"github.com/19chonm/461_1_23/fileio"
)

func runTask(url string, ratingch chan<- fileio.Rating) {
	fmt.Println("My job is", url)
	r := fileio.Rating{NetScore: rand.Float64(), Url: url} // Placeholder until real data implemented
	ratingch <- r
}
