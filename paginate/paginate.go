package paginate

import (
	"context"
	"time"
	"github.com/google/go-github/github"
)

type Paginate interface {
	IssuesCreated(ctx context.Context, client *github.Client, repoOwner string, repoName string,
				  username string, yearAgo time.Time)
	Pulls(ctx context.Context, client *github.Client, repoOwner string, repoName string,
		  username string, yearAgo time.Time) ([]*github.PullRequest, error)
	Commits(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, yearAgo time.Time, 
			repo *github.Repository)
}

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// returns all issues created by the user in that repo
func IssuesCreated(ctx context.Context, client *github.Client, repoOwner string, repoName string,
				   username string, yearAgo time.Time) ([]*github.Issue, error) {
	var l []*github.Issue
	opt := &github.IssueListByRepoOptions{Creator: username, Since: yearAgo, State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.Issues.ListByRepo(ctx, repoOwner, repoName, opt)
		if err != nil {
			return nil, err
		}
		l = append(l, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return l, nil
}

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// returns all pull requests for that repo
func Pulls(ctx context.Context, client *github.Client, repoOwner string, repoName string,
		   username string, yearAgo time.Time) ([]*github.PullRequest, error) {
	var l []*github.PullRequest
	opt := &github.PullRequestListOptions{State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.PullRequests.List(ctx, repoOwner, repoName, opt)
		if err != nil {
			return nil, err
		}
		l = append(l, list...)
		if resp.NextPage == 0 || opt.Page == 10 {
			break
		}
		opt.Page = resp.NextPage
	}
	return l, nil
}

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// and a repo, returns all commits for that repo
func Commits(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, yearAgo time.Time, 
			 repo *github.Repository) ([]*github.RepositoryCommit, error) {
	var commitsList []*github.RepositoryCommit
		opt := &github.CommitsListOptions{SHA: "master", Author: username, Since: yearAgo, ListOptions: github.ListOptions{PerPage: 30}}
		if repo.GetSize() != 0 {
		for {
			list, resp, err := client.Repositories.ListCommits(ctx, repoOwner, repoName, opt)
			if err != nil {
				return nil, err
			}
			commitsList = append(commitsList, list...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return commitsList, nil
}
