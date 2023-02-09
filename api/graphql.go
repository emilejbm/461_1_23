/*
Gets factors that are used to build the rating for correctness of modules via
GitHub's GraphQL API. getCorrectnessFactors function returns factors in the
order of watchers, stargazers, totalCommits. A data type in the same form as
the query structure is required to convert string to json. From json, the data
is returned.
*/

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CorrectnessFactors struct {
	Data struct {
		Repository struct {
			StargazerCount int64
			Watchers       struct {
				TotalCount int64
			}
			DefaultBranchRef struct {
				Target struct {
					History struct {
						TotalCount int64
					}
				}
			}
		}
	}
}

func buildQuery(ownerName string, repoName string) (query map[string]string) {
	var correctnessQuery = map[string]string{
		"query": `
		{
			repository(owner:` + `"` + ownerName + `", name:` + `"` + repoName + `") {
				name 
				stargazerCount
				watchers {
					totalCount
				}
				defaultBranchRef {
					target {
						... on Commit {
							history(first:0) {
								totalCount
							}
						}
					}
				}
			}
		}`,
	}

	return correctnessQuery
}

func getGraphqlResponse(ownerName string, repoName string) []uint8 {

	query := buildQuery(ownerName, repoName)

	jsonValue, _ := json.Marshal(query)
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+"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")
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Reading body failed with error %s\n", err)
	}

	return data
}

func getCorrectnessFactors(ownerName string, repoName string) (watchers int64, stargazers int64, totalCommits int64) {

	ownerName := "expressjs"
	repoName := "express"

	data := getGraphqlResponse(ownerName, repoName)
	var factors CorrectnessFactors
	newerr := json.Unmarshal([]byte(string(data)), &factors)

	if newerr != nil {
		fmt.Printf("error %s\n", newerr)
	}

	watchers := factors.Data.Repository.StargazerCount
	stargazers := factors.Data.Repository.Watchers.TotalCount
	totalCommits := factors.Data.Repository.DefaultBranchRef.Target.History.TotalCount

	return watchers, stargazers, totalCommits

}
