package main

import (
	"context"
	"flag"
	"fmt"
	"os"
    	"sync"
	"time"
	"issues"
	"pulls"
	"commits"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"paginate"
	"util"
)

func main() {
	var wg sync.WaitGroup

	ctx, client := authentication()
	var yearAgo = time.Now().AddDate(-1, 0, 0)

	flag.Parse() 
	args := flag.Args()
	username := args[0]
	if len(args) < 1 {
		fmt.Println("go run fetch <username>")
		os.Exit(1)
	}

	mIssues := make(map[int]int)
    	mPulls := make(map[int]int)
	mCommits := make(map[int]int)
	repoList := paginate.Repo(ctx, client, username)
	for i := 0; i < len(repoList); i++ {
		repo := repoList[i]
		repoName := repo.GetName()
		repoOwner := repo.GetOwner().GetLogin()
		issueListCreator := paginate.IssuesCreated(ctx, client, repoOwner, repoName, username, yearAgo)
		issueList := paginate.IssueEvents(ctx, client, repoOwner, repoName, username, yearAgo)
		pullsList := paginate.Pulls(ctx, client, repoOwner, repoName, username, yearAgo)
		commitsList := paginate.Commits(ctx, client, repoOwner, repoName, username, yearAgo, repo)
		wg.Add(6)
		go func() {
			commits.GetCommitTimes(commitsList, nil, mCommits)
			wg.Done()
		}()
		go func() {
			pulls.GetPullsCreatedTimes(ctx, client, repoOwner, repoName, username, pullsList, mPulls)
			wg.Done()
		}()
		go func() {
			pulls.GetPullsEventTimes(ctx, client, repoOwner, repoName, username, pullsList, mPulls)
			wg.Done()
		}()
		go func() {
			pulls.GetPullsReviewRequestTimes(ctx, client, repoOwner, repoName, username, pullsList, mPulls)
			wg.Done()
		}()
		go func() {
			issues.GetIssueCreatedTimes(issueListCreator, mIssues)
			wg.Done()
		}()
		go func() {
			issues.GetIssueEventTimes(ctx, client, repoOwner, repoName, username, issueList, mIssues)
			wg.Done()
		}()
	}
	result := util.ComputeContr(mIssues, mPulls, mCommits)
	fmt.Println(result)
}

func throwError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func authentication() (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}
