package pulls_test

import(
	"testing"
	"context"
	"reflect"
	"net/http"
	"net/http/httptest"
	"github.com/google/go-github/github"
	"net/url"
	"fmt"
)

func TestGetPullsCreatedTimes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/pulls/1/reviews", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" {
			t.Error()
		}
		want := url.Values{}
		want.Set("page", "2")

		r.ParseForm()
		if !reflect.DeepEqual(r.Form, want) {
			t.Errorf("Request parameters: %v, want %v", r.Form, want)
		}
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})
	var id1 int64 = 1
	var id2 int64 = 2
	opt := &github.ListOptions{Page: 2}
	reviews, _, err := client.PullRequests.ListReviews(context.Background(), "o", "r", 1, opt)
	if err != nil {
		t.Errorf("PullRequests.ListReviews returned error: %v", err)
	}
	wantReview := []*github.PullRequestReview{
		{ID: &id1},
		{ID: &id2},
	}
	time := reviews[0].GetSubmittedAt().Format("2006-01-02T15:04:05Z07:00")
	wantTime := wantReview[0].GetSubmittedAt().Format("2006-01-02T15:04:05Z07:00")
	if !reflect.DeepEqual(time, wantTime) {
		t.Errorf("GetIssuesCreated returned %+v, want %+v", time, wantTime)
	}
}

func setup() (client *github.Client, mux *http.ServeMux, teardown func()) {
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