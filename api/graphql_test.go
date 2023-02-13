package api

import (
	"testing"
)

/*
Stats for number of watchers / stargazers / totalCommits (as of feb 11, 11:02 am)
Will be reduced by 5% to account for volatility
facebook/react:			201883 / 6640 / 15506
axios/axios:			98724 / 1207 / 1426
expressjs/express:		59880 / 1729 / 5756
request/request:		25607 / 441 / 2270
nullivex/nodist: 		1470 / 59 / 415
*/

func TestCorrectFactors(t *testing.T) {
	reactFactors := [3]int64{191788, 6308, 14730}
	axiosFactors := [3]int64{93787, 1146, 1354}
	expressFactors := [3]int64{56886, 1642, 5468}
	requestFactors := [3]int64{24326, 418, 2156}
	nodistFactors := [3]int64{1396, 56, 394}

	owner_repo_names := [5]string{
		"https://github.com/facebook/react", "https://github.com/axios/axios",
		"https://github.com/expressjs/express", "https://github.com/request/request",
		"https://github.com/nullivex/nodist",
	}

	var factors [][3]int64
	factors = append(factors, reactFactors)
	factors = append(factors, axiosFactors)
	factors = append(factors, expressFactors)
	factors = append(factors, requestFactors)
	factors = append(factors, nodistFactors)

	for i := range owner_repo_names {
		f1, f2, f3, err := GetCorrectnessFactors(owner_repo_names[i])

		if int64(f1) < (factors[i][0]) || int64(f2) < factors[i][1] || int64(f3) < factors[i][2] {
			t.Errorf("user got: %d, %d, %d. want at least: %d, %d, %d", f1, f2, f3, factors[i][0], factors[i][1], factors[i][2])
		}
		if err != nil {
			t.Errorf("got error %s", err)
		}
	}
}

// Checking is done so that a valid response is received. Not for correctness

func TestGoodGithubNames(t *testing.T) {
	goodUrl := "https://github.com/facebook/react"
	m1, m2, m3, err := GetCorrectnessFactors(goodUrl)
	if m1 == 0 && m2 == 0 && m3 == 0 && err != nil {
		t.Errorf("user got error %s", err)
	}

}

func TestBadGithubNames(t *testing.T) {
	badUrl := "xjaop!@#$%^&asd\naliru"
	m1, m2, m3, err := GetCorrectnessFactors(badUrl)
	if err == nil || m1 != 0 || m2 != 0 || m3 != 0 {
		t.Errorf("expected error, got %d, %d, %d", m1, m2, m3)
	}

}
