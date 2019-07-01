package paginate 

import (
	"github.com/google/go-github/github"
	"context"
	"time"
	"os"
)

func Repo(ctx context.Context, client *github.Client, username string) ([]*github.Repository) {
	var repoList []*github.Repository
	optRepo := &github.RepositoryListOptions{Type: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		repos, resp, err := client.Repositories.List(ctx, username, optRepo)
		if err != nil {
			os.Exit(1)
		}
		repoList = append(repoList, repos...)
		if resp.NextPage == 0 {
			break
		}
		optRepo.Page = resp.NextPage
	}
	return repoList
}

func IssuesCreated(ctx context.Context, client *github.Client, repoOwner string, repoName string, 
				   username string, yearAgo time.Time) ([]*github.Issue)  {
	var issueListCreator []*github.Issue
	opt := &github.IssueListByRepoOptions{Creator: username, Since: yearAgo, State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.Issues.ListByRepo(ctx, repoOwner, repoName, opt)
		if err != nil {
			os.Exit(1) 
		}
		issueListCreator = append(issueListCreator, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return issueListCreator
}

func IssueEvents(ctx context.Context, client *github.Client, repoOwner string, repoName string, 
	username string, yearAgo time.Time) ([]*github.Issue) {
	var issueList []*github.Issue
	opt := &github.IssueListByRepoOptions{Since: yearAgo, State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.Issues.ListByRepo(ctx, repoOwner, repoName, opt)
		if err != nil {
			os.Exit(1) 
		}
		issueList = append(issueList, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return issueList
} 

func Pulls(ctx context.Context, client *github.Client, repoOwner string, repoName string, 
		   username string, yearAgo time.Time) ([]*github.PullRequest) {
	var pullsList []*github.PullRequest
	opt := &github.PullRequestListOptions{State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.PullRequests.List(ctx, repoOwner, repoName, opt)
		if err != nil {
			os.Exit(1) 
		}
		pullsList = append(pullsList, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return pullsList
}

func Commits(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, 
			 yearAgo time.Time, repo *github.Repository) ([]*github.RepositoryCommit) {
	var commitsList []*github.RepositoryCommit
	opt := &github.CommitsListOptions{Author: username, Since: yearAgo, ListOptions: github.ListOptions{PerPage: 30},}
	if repo.GetSize() != 0 {
		for {
			list, resp, err := client.Repositories.ListCommits(ctx, repoOwner, repoName, opt)
			if err != nil {
				os.Exit(1)
			}
			commitsList = append(commitsList, list...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return commitsList
}