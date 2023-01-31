# 461_1_23
Mimi Chon
Anna Shen
Emile Baez
Ben Schwartz

# Helpful Commands
go run .
go build .


# Basic API curl for responsiveness
curl   -H "Accept: application/vnd.github+json"   -H "Authorization: Bearer <token>"  -H "X-GitHub-Api-Version: 2022-11-28"   https://api.github.com/repos/19chonm/461_1_23/issues?state=closed

Use the following metrics to calculate days between:
    "state": "closed",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2023-01-22T22:28:21Z",
    "updated_at": "2023-01-22T22:28:50Z",
    "closed_at": "2023-01-22T22:28:44Z",

https://stackoverflow.com/questions/47063026/unable-to-see-the-closed-issues-with-the-github-api
https://stackoverflow.com/questions/58665002/using-github-list-issues-for-a-repository-api


### Basic API curl for looking at contributors
curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer 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"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/19chonm/461_1_23/stats/contributors

https://docs.github.com/en/rest/metrics/statistics?apiVersion=2022-11-28#get-all-contributor-commit-activity