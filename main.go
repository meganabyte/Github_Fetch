package main

import (
	"commits"
	"context"
	"flag"
	"fmt"
	"issues"
	"os"
	"paginate"
	"pulls"
	"time"
	"util"
	"repos"
	"sync"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	var yearAgo = time.Now().AddDate(-1, 0, 0)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	flag.Parse()
	args := flag.Args()
	username := args[0]
	if len(args) < 2 {
		fmt.Println("go run main <username> <OAUTH token>")
		os.Exit(1)
	} 
	token := args[1]
	ctx, client := authentication(token)

	mIssues := make(map[int]int)
	mPulls := make(map[int]int)
	mCommits := make(map[int]int)
	repoList := repos.GetRepoList(ctx, client, username)
	commitRepoList, _ := repos.GetCommitRepoList(ctx, client, repoList)

	for _, repo := range commitRepoList {
		repoName, repoOwner := repos.GetRepoInfo(repo)
		var list1 []*github.Issue
		var list2 []*github.PullRequest
		var list3 []*github.RepositoryCommit
		wg.Add(3)
		go func() {
			list1, _ = paginate.IssuesCreated(ctx, client, repoOwner, repoName, username, yearAgo)
			wg.Done()
		}()
		go func() {
			list2, _ = paginate.Pulls(ctx, client, repoOwner, repoName, username, yearAgo)
			wg.Done()
		}()
		go func() {
			list3, _ = paginate.Commits(ctx, client, repoOwner, repoName, username, yearAgo, repo)
			wg.Done()
		}()
		wg.Wait()
		wg2.Add(3)
		go func(list1 []*github.Issue) {
			issues.GetIssueCreatedTimes(list1, mIssues)
			wg2.Done()
		}(list1)
		go func(list2 []*github.PullRequest) {
			pulls.GetPullsReviewRequestTimes(ctx, client, repoOwner, repoName, username, list2, mPulls)
			wg2.Done()
		}(list2)
		go func(list3 []*github.RepositoryCommit) {
			commits.GetCommitTimes(list3, mCommits)
			wg2.Done()
		}(list3)
		wg2.Wait()
	}
	result := util.ComputeContr(mIssues, mPulls, mCommits)
	fmt.Println(result)
}

// given a context and a reference to a github client, creates an authenticated
// github client
func authentication(token string) (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}
