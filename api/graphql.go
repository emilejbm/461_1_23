/*
Gets factors that are used to build the rating for correctness of modules via
GitHub's GraphQL API. getCorrectnessFactors function returns factors in the
order of watchers, stargazers, totalCommits. A data type in the same form as
the query structure is required to convert string to json. From json, the data
is returned.

--Use of GitHub token needs to be changed
*/

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func buildCorrectnessQuery(ownerName string, repoName string) (query map[string]string) {
	var correctnessQuery = map[string]string{
		"query": `
		{
			repository(owner:` + `"` + ownerName + `", name:` + `"` + repoName + `") { 
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

func getCorrectnessFactors(ownerName string, repoName string) (watchers int64, stargazers int64, totalCommits int64, err error) {

	query := buildCorrectnessQuery(ownerName, repoName)
	token, ok := os.LookupEnv("GITHUB_TOKEN")

	if !ok {
		return 0, 0, 0, fmt.Errorf("validateInput: Error getting token from environment variable")
	}

	jsonValue, _ := json.Marshal(query)
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return 0, 0, 0, fmt.Errorf("The GraphQL query failed with error %s\n", err)
	}

	var factors CorrectnessFactors
	err = json.NewDecoder(res.Body).Decode(&factors)

	if err != nil {
		return 0, 0, 0, fmt.Errorf("Reading body failed with errorr %s\n", err)
	}

	watchers = factors.Data.Repository.StargazerCount
	stargazers = factors.Data.Repository.Watchers.TotalCount
	totalCommits = factors.Data.Repository.DefaultBranchRef.Target.History.TotalCount

	return watchers, stargazers, totalCommits, nil
}
