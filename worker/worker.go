package worker

import (
	"fmt"

	"github.com/19chonm/461_1_23/api"
	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/metrics"
)

func runTask(url string, ratingch chan<- fileio.Rating) {
	fmt.Println("My job is", url)

	// Get Data from Github API
	license_key, err := api.GetRepoLicense(url)
	if err != nil {
		fmt.Println("worker: ERROR Unable to get data for ", url, " License Errored:", err)
		return
	}

	avg_lifespan, err := api.GetRepoIssueAverageLifespan(url)
	if err != nil {
		fmt.Println("worker: ERROR Unable to get data for ", url, " AvgLifespan Errored:", err)
		return
	}

	top_recent_commits, total_recent_commits, err := api.GetRepoContributors(url)
	if err != nil {
		fmt.Println("worker: ERROR Unable to get data for ", url, " ContributorsCommits Errored:", err)
		return
	}

	// Download repository and scan
	rampup_score, err := metrics.ScanRepo(url)
	if err != nil {
		fmt.Println("worker: ERROR Unable to get data for ", url, " ScanRepo Errored:", err)
		return
	}

	// Compute scores
	correctness_score := metrics.ComputeCorrectness(0, 0, 0) // no data yet
	responsiveness_score := metrics.ComputeResponsiveness(avg_lifespan)
	busfactor_score := metrics.ComputeBusFactor(top_recent_commits, total_recent_commits)
	license_score := metrics.ComputeLicenseScore(license_key)

	rampup_factor := metrics.Factor{Weight: 0.15, Value: rampup_score, AllOrNothing: false}
	correctness_factor := metrics.Factor{Weight: 0.15, Value: correctness_score, AllOrNothing: false}
	responsiveness_factor := metrics.Factor{Weight: 0.4, Value: responsiveness_score, AllOrNothing: false}
	busfactor_factor := metrics.Factor{Weight: 0.3, Value: busfactor_score, AllOrNothing: false}
	license_factor := metrics.Factor{Weight: 1.0, Value: float64(license_score), AllOrNothing: true}

	// Produce final rating
	factors := []metrics.Factor{rampup_factor, correctness_factor, responsiveness_factor, busfactor_factor, license_factor}
	r := fileio.Rating{NetScore: metrics.ComputeNetScore(factors),
		Rampup:         rampup_score,
		Url:            url,
		License:        float64(license_score),
		Busfactor:      busfactor_score,
		Responsiveness: responsiveness_score,
		Correctness:    correctness_score,
	}
	ratingch <- r // Send rating to rating channel to be sorted
}
