package api

import (
	"fmt"
	"io"
	// "log"
	"net/http"
	"os"
	"testing"
	"bytes"
	"net/url"
	// "time"
	// "encoding/json"
)

type TestType struct {
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}
// Input URL Tests
func Test_ValidateInput_Success(t *testing.T) {
	var goodInputUrl string = "https://github.com/facebook/react"
	var correctUser string = "facebook"
	var correctRepo string = "react"
	correctToken, _ := os.LookupEnv("GITHUB_TOKEN")
	user, repo, token, ok := ValidateInput(goodInputUrl)
	if user != correctUser {
		t.Errorf("user got: %s, want: %s.", user, correctUser)
	}
	if repo != correctRepo {
		t.Errorf("repo got: %s, want: %s.", repo, correctRepo)
	}
	if token != correctToken {
		t.Errorf("token got: %s, want: %s.", token, correctToken)
	}
	if ok != nil {
		t.Errorf("ok was not nil: %s", ok.Error())
	}
}

func Test_ValidateInput_BadURL(t *testing.T) {
	badUrls := [...]string{"https://google.com/facebook/react", "https://github.com/someuser", "\n", ""}
	var badUser string = ""
	var badRepo string = ""
	var badToken string = ""
	badOk := fmt.Errorf("some api error")

	for _, badUrl := range badUrls {
		user, repo, token, ok := ValidateInput(badUrl)
		if user != badUser {
			t.Errorf("user got: %s, want: %s.", user, badUser)
		}
		if repo != badRepo {
			t.Errorf("repo got: %s, want: %s.", repo, badRepo)
		}
		if token != badToken {
			t.Errorf("token got: %s, want: %s.", token, badToken)
		}
		if ok == nil {
			t.Errorf("ok got: %s, want: %s.", ok, badOk.Error())
		}
	}
}

func Test_ValidateInput_NoToken(t *testing.T) {
	var inputUrl string = "https://github.com/facebook/react"
	var badUser string = ""
	var badRepo string = ""
	var badToken string = ""

	t.Setenv("GITHUB_TOKEN", "") // make t restore GITHUB_TOKEN on cleanup
	os.Unsetenv("GITHUB_TOKEN")
	user, repo, token, ok := ValidateInput(inputUrl)
	if user != badUser {
		t.Errorf("user got: %s, want: %s.", user, badUser)
	}
	if repo != badRepo {
		t.Errorf("repo got: %s, want: %s.", repo, badRepo)
	}
	if token != badToken {
		t.Errorf("token got: %s, want: %s.", token, badToken)
	}
	if ok == nil {
		t.Errorf("ok was nil, expected error")
	}
}

func Test_ValidateInput_EmptyToken(t *testing.T) {
	var inputUrl string = "https://github.com/facebook/react"
	var badUser string = ""
	var badRepo string = ""
	var badToken string = ""

	t.Setenv("GITHUB_TOKEN", "")
	user, repo, token, ok := ValidateInput(inputUrl)
	if user != badUser {
		t.Errorf("user got: %s, want: %s.", user, badUser)
	}
	if repo != badRepo {
		t.Errorf("repo got: %s, want: %s.", repo, badRepo)
	}
	if token != badToken {
		t.Errorf("token got: %s, want: %s.", token, badToken)
	}
	if ok == nil {
		t.Errorf("ok was nil, expected error")
	}
}

func Test_ValidateInput_EmptyURL(t *testing.T) {
	var inputUrl string = ""
	_, _, _, ok := ValidateInput(inputUrl)
	if ok == nil {
		t.Errorf("ok was nil, expected error")
	}
}

func Test_DecodeResponse_Success(t *testing.T) {
	res := http.Response{
		Body: io.NopCloser(bytes.NewBufferString("{\"foo\": 461, \"bar\": \"Project\"}")),
	}
	correctFoo := 461
	correctBar := "Project"
	jsonRes, err := DecodeResponse[TestType](&res)

	if jsonRes.Foo != correctFoo {
		t.Errorf("jsonRes.Foo got: %d want: %d.", jsonRes.Foo, correctFoo)
	}
	if jsonRes.Bar != correctBar {
		t.Errorf("jsonRes.Bar got: %s want: %s.", jsonRes.Bar, correctBar)
	}
	if err != nil {
		t.Errorf("err got: %v", err)
	}
}

func Test_DecodeResponse_Failure(t *testing.T) {
	res := http.Response{
		Body: io.NopCloser(bytes.NewBufferString("this isn't json")),
	}
	_, err := DecodeResponse[TestType](&res)

	if err == nil {
		t.Errorf("err was nil")
	}
}

func Test_SendGithubRequestHelper_Success(t *testing.T) {
	endpoint := "https://api.github.com/users/octocat/orgs"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	res, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if res == nil {
		t.Errorf("res is nil")
	}

	if err != nil {
		t.Errorf("Got unexpected error %s", err.Error())
	} 
	if (statusCode != 200) {
		t.Errorf("GitHub request responded with error code %d", statusCode)
	}
}

func Test_SendGithubRequestHelper_BadEndpoint(t *testing.T) {
	endpoint := "https://api.github.com/bad_endpoint"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode == 200) {
		t.Errorf("Got 200 status code")
	}
}

func Test_SendGithubRequestHelper_NotFound(t *testing.T) {
	endpoint := "https://api.github.com/repos/fakeuser/fakerepohopefully"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode != 404) {
		t.Errorf("want 404 status code, got %d", statusCode)
	}
}

func Test_SendGithubRequestHelper_BadToken(t *testing.T) {
	endpoint := "https://api.github.com/user/octocat/orgs"
	token := "invalid_token"

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode == 200) {
		t.Errorf("Got 200 status code")
	}
}

type EmptyResponse struct {

}
func (self EmptyResponse) Validate() bool {
	return true;
}

type InvalidResponse struct {

}
func (self InvalidResponse) Validate() bool {
	return false;
}

func Test_SendGithubRequest_UnexpectedPagination(t *testing.T) {
	endpoint := "https://api.github.com/repos/octocat/Spoon-Knife/issues"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequest[EmptyResponse](endpoint, token)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Did not expect pagination" {
		t.Errorf("Expected error \"Did not expect pagination\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

func Test_SendGithubRequest_ValidationFailure(t *testing.T) {
	endpoint := "https://api.github.com/repos/octocat/Spoon-Knife"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequest[InvalidResponse](endpoint, token)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Failed to parse GitHub response" {
		t.Errorf("Expected error \"Failed to parse GitHub response\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

func Test_SendGithubRequest_RequestFailure(t *testing.T) {
	endpoint := "bad_endpoint"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequest[EmptyResponse](endpoint, token)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Failed to send HTTP request" {
		t.Errorf("Expected error \"Failed to send HTTP request\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

// func Test_SendGithubRequest_202Loop(t *testing.T) {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// w.Write()
// 	})
// 	go func() {
// 		log.Fatal(http.ListenAndServe(":3461", nil))
// 	}()
// 	time.Sleep(1 * time.Second)
// 	endpoint := "http://localhost:3461"
// 	token := os.Getenv("GITHUB_TOKEN")
// 	_, err, statusCode := SendGithubRequest[EmptyResponse](endpoint, token)
// 	if err == nil {
// 		t.Errorf("err was nil")
// 	}
// 	if err.Error() != "Github request exceed max retry count for error code 202" {
// 		t.Errorf("got wrong err %s", err.Error())
// 	}
// 	if statusCode != 500 {
// 		t.Errorf("Expected statusCode 500, got %d", statusCode)
// 	}
// }

func Test_SendGithubRequestList_ValidationFailure(t *testing.T) {
	endpoint := "https://api.github.com/repos/octocat/Spoon-Knife/issues"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequestList[InvalidResponse](endpoint, token, 1)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Failed to parse GitHub response" {
		t.Errorf("Expected error \"Failed to parse GitHub response\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

func Test_SendGithubRequestList_RequestFailure(t *testing.T) {
	endpoint := "bad_endpoint"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequestList[EmptyResponse](endpoint, token, 1)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Failed to send HTTP request" {
		t.Errorf("Expected error \"Failed to send HTTP request\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

func Test_SendGithubRequestList_ParamFailure(t *testing.T) {
	endpoint := "\n"
	token := os.Getenv("GITHUB_TOKEN")

	_, err, statusCode := SendGithubRequestList[EmptyResponse](endpoint, token, 1)
	if err == nil {
		t.Errorf("err was nil")
	}
	if err.Error() != "Failed to set query parameter" {
		t.Errorf("Expected error \"Failed to set query parameter\", got \"%s\"", err.Error())
	}
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got %d", statusCode)
	}
}

func Test_SetQueryParameter_SuccessNotExists(t *testing.T) {
	endpoint := "https://example.com/path?foo=bar"
	SetQueryParameter(&endpoint, "baz", "quux")
	urlObject, _ := url.Parse(endpoint)
	query := urlObject.Query()
	
	if query.Get("baz") != "quux" {
		t.Errorf("want baz to be quux, got %s", query.Get("baz"))
	}
	if query.Get("foo") != "bar" {
		t.Errorf("want foo to be bar, got %s", query.Get("foo"))
	}
}

func Test_SetQueryParameter_SuccessExists(t *testing.T) {
	endpoint := "https://example.com/path?foo=bar"
	SetQueryParameter(&endpoint, "foo", "baz")
	urlObject, _ := url.Parse(endpoint)
	query := urlObject.Query()

	if query.Get("foo") != "baz" {
		t.Errorf("want foo to be bar, got %s", query.Get("foo"))
	}
}

func Test_SetQueryParameter_Error(t *testing.T) {
	endpoint := "\n" // bad endpoint
	SetQueryParameter(&endpoint, "foo", "bar")
	_, err := url.Parse(endpoint)

	if err == nil {
		t.Errorf("err is nil")
	}
}

func Test_GetRepoLicense_Success(t *testing.T) {
	goodInputUrl := "https://github.com/octocat/git-consortium" // FIX: this is not a 
	license, err := GetRepoLicense(goodInputUrl) 
	
	if !(license == "mit") {
		t.Errorf("Expected license to be mit, got %s", license)
	}
	if err != nil {
		t.Errorf("Got RepoLicense err")
	}
}

func Test_GetRepoLicense_SuccessNoLicense(t *testing.T) {
	goodInputUrl := "https://github.com/octocat/Spoon-Knife" // FIX: this is not a 
	license, err := GetRepoLicense(goodInputUrl) 
	
	if !(license == "") {
		t.Errorf("Expected license to be empty, got %s", license)
	}
	if err != nil {
		t.Errorf("Got RepoLicense err: %s", err.Error())
	}
}

//TODO: Add error case as well
func Test_GetRepoContributors_Success(t *testing.T) {
	goodInputUrl := "https://github.com/facebook/react" // FIX: this is not a 
	top3, total, err := GetRepoContributors(goodInputUrl) 
	if !(top3 > 0) {
		t.Errorf("RepoContributors.top3 want d, got: d")
	}
	if !(total > 0) {
		t.Errorf("RepoContributors.total want d, got: d")
	}
	if err != nil {
		t.Errorf("Got RepoContributors err")
	}
}

//TODO: Add error case as well
func Test_GetRepoIssueAverageLifespan_Success(t *testing.T) {
	goodInputUrl := "https://github.com/facebook/react" // FIX: this is not a 
	avgLifespan, err := GetRepoIssueAverageLifespan(goodInputUrl) 
	if !(avgLifespan > 0) {
		t.Errorf("IssueAvgLifesapn.avgLifespan want d, got: d")
	}
	if err != nil {
		t.Errorf("Got IssueAvgLifesapn err")
	}
}
