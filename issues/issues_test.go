package issues_test 

import(
	"testing"
	"context"
	"reflect"
	"net/http"
	"net/http/httptest"
	"github.com/google/go-github/github"
	"net/url"
	"time"
	"fmt"
)

func TestGetIssueCreatedTimes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/issues", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" {
			t.Error()
		}
		want := url.Values{}
		want.Set("milestone", "*")
		want.Set("state", "closed")
		want.Set("assignee", "a")
		want.Set("creator", "c")
		want.Set("mentioned", "m")
		want.Set("labels", "a,b")
		want.Set("sort", "updated")
		want.Set("direction", "asc")
		want.Set("since", "2013-08-01T00:00:00Z")

		r.ParseForm()
		if !reflect.DeepEqual(r.Form, want) {
			t.Errorf("Request parameters: %v, want %v", r.Form, url.Values{})
		}
		fmt.Fprintf(w, `[{"number" : 1}]`)
	 }) 
	 opt := &github.IssueListByRepoOptions{
		"*", "closed", "a", "c", "m", []string{"a", "b"}, "updated", "asc", 
		time.Date(2013, time.August, 1, 0, 0, 0, 0, time.UTC), 
		github.ListOptions{0,0},
	}
	issues, _, err := client.Issues.ListByRepo(context.Background(), "o", "r", opt)
	if err != nil {
		t.Errorf("GetIssuesCreated returned error %v", err)
	}
	num := 1
	time := issues[0].GetCreatedAt().Format("2006-01-02T15:04:05Z07:00")
	wantIssues := []*github.Issue{{Number: &num}}
	wantTime := wantIssues[0].GetCreatedAt().Format("2006-01-02T15:04:05Z07:00")
	if !reflect.DeepEqual(time, wantTime) {
		t.Errorf("GetIssuesCreated returned %+v, want %+v", time, wantTime)
	}
}


func setup() (client *github.Client, mux *http.ServeMux, teardown func()){
	mux = http.NewServeMux() // http multiplexer used with test server 
	handler := http.NewServeMux()
	handler.Handle("/api-v3/", http.StripPrefix("/api-v3", mux))
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){})
	server := httptest.NewServer(handler) // test HTTP server to mock API responses

	// new github client to user for testing, configured to user test server
	client = github.NewClient(nil)
	url, _ := url.Parse(server.URL + "/api-v3/")
	client.BaseURL = url
	client.UploadURL = url

	return client, mux, server.Close
}