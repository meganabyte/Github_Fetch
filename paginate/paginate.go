package paginate

import (
	"context"
	"time"
	"util"
	"github.com/google/go-github/github"
)

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// returns a list of all issues created by the user in that repo
func IssuesCreated(ctx context.Context, client *github.Client, repoOwner string, repoName string,
				   username string, yearAgo time.Time) []*github.Issue {
	var issueListCreator []*github.Issue
	opt := &github.IssueListByRepoOptions{Creator: username, Since: yearAgo, State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.Issues.ListByRepo(ctx, repoOwner, repoName, opt)
		util.ThrowError(err)
		issueListCreator = append(issueListCreator, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return issueListCreator
}

// given a context, a client, a repo's owner, a repo's name, and a user's login, and the time a year ago,
// returns a list of all issues in that repo
func IssueEvents(ctx context.Context, client *github.Client, repoOwner string, repoName string,
				 username string, yearAgo time.Time) []*github.Issue {
	var issueList []*github.Issue
	opt := &github.IssueListByRepoOptions{Since: yearAgo, State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.Issues.ListByRepo(ctx, repoOwner, repoName, opt)
		util.ThrowError(err)
		issueList = append(issueList, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return issueList
}

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// returns a list of all pull requests for that repo
func Pulls(ctx context.Context, client *github.Client, repoOwner string, repoName string,
		   username string, yearAgo time.Time) []*github.PullRequest {
	var pullsList []*github.PullRequest
	opt := &github.PullRequestListOptions{State: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		list, resp, err := client.PullRequests.List(ctx, repoOwner, repoName, opt)
		util.ThrowError(err)
		pullsList = append(pullsList, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return pullsList
}

// given a context, a client, a repo's owner, a repo's name, a user's login, and the time a year ago,
// and a repo, a returns a list of all commits for that repo
func Commits(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
			 yearAgo time.Time, repo *github.Repository) []*github.RepositoryCommit {
	var commitsList []*github.RepositoryCommit
	opt := &github.CommitsListOptions{Author: username, Since: yearAgo, ListOptions: github.ListOptions{PerPage: 30}}
	if repo.GetSize() != 0 {
		for {
			list, resp, err := client.Repositories.ListCommits(ctx, repoOwner, repoName, opt)
			util.ThrowError(err)
			commitsList = append(commitsList, list...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return commitsList
}
