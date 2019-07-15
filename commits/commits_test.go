package commits_test

import (
	"testing"
	"context"
	"reflect"
	"net/http"
	"net/http/httptest"
	"github.com/google/go-github/github"
	"net/url"
	"time"
	"fmt"
	"os"

	
)

func TestGetCommitTimes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/commits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Error()
		}
		want := url.Values{}
		want.Set("sha", "s")
		want.Set("path", "p")
		want.Set("author", "a")
		want.Set("since", "2013-08-01T00:00:00Z")
		want.Set("until", "2013-09-03T00:00:00Z")

		r.ParseForm()
		if !reflect.DeepEqual(r.Form, want) {
			t.Errorf("Request parameters: %v, want %v", r.Form, url.Values{})
		}
		fmt.Fprintf(w, `[{"sha" : "s"}]`)
	 }) 
	 opt := &github.CommitsListOptions{
		SHA:    "s",
		Path:   "p",
		Author: "a",
		Since:  time.Date(2013, time.August, 1, 0, 0, 0, 0, time.UTC),
		Until:  time.Date(2013, time.September, 3, 0, 0, 0, 0, time.UTC),
	}
	commits, _, err := client.Repositories.ListCommits(context.Background(), "o", "r", opt)
	if err != nil {
		t.Errorf("Commits returned error %v", err)
	}
	time := commits[0].Commit.GetAuthor().GetDate().Format("2006-01-02T15:04:05Z07:00")
	wantCommits := []*github.RepositoryCommit{{SHA: &opt.SHA}}
	wantTime := wantCommits[0].Commit.GetAuthor().GetDate().Format("2006-01-02T15:04:05Z07:00")
	if !reflect.DeepEqual(time, wantTime) {
		t.Errorf("GetCommitTimes returned %+v, want %+v", time, wantTime)
	}
}

func setup() (client *github.Client, mux *http.ServeMux, teardown func()){
	mux = http.NewServeMux() // http multiplexer used with test server 
	handler := http.NewServeMux()
	handler.Handle("/api-v3/", http.StripPrefix("/api-v3", mux))
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(os.Stderr)
	})
	server := httptest.NewServer(handler) // test HTTP server to mock API responses

	// new github client to user for testing, configured to user test server
	client = github.NewClient(nil)
	url, _ := url.Parse(server.URL + "/api-v3/")
	client.BaseURL = url
	client.UploadURL = url

	return client, mux, server.Close
}
