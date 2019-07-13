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
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	var yearAgo = time.Now().AddDate(-1, 0, 0)

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

	for _, repo := range repoList {
		repoName, repoOwner := repos.GetRepoInfo(repo)
		fmt.Println(repoName, repoOwner)
		c1, _ := paginate.Commits(ctx, client, repoOwner, repoName, username, yearAgo, repo)
		c2, _ := paginate.IssuesCreated(ctx, client, repoOwner, repoName, username, yearAgo)
		c3, _ := paginate.IssueEvents(ctx, client, repoOwner, repoName, username, yearAgo)
		c4, _ := paginate.Pulls(ctx, client, repoOwner, repoName, username, yearAgo)
		commits.GetCommitTimes(c1, mCommits)
		issues.GetIssueEventTimes(ctx, client, repoOwner, repoName, username, c3, mIssues)
		issues.GetIssueCreatedTimes(c2, mIssues)
		pulls.GetPullsReviewRequestTimes(ctx, client, repoOwner, repoName, username, c4, mPulls)
		time.Sleep(10 * time.Second)
	}
	fmt.Println("Issues", mIssues)
	fmt.Println("Pulls", mPulls)
	fmt.Println("Commits", mCommits)
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
