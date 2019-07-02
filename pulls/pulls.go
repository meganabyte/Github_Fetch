package pulls

import (
	"util"
	"context"
	"github.com/google/go-github/github"
)

// given a context, a client, a repo's owner, a repo's name, a user's login, a list of pull requests for a repo
// and a record of pull requests for a user, adds a mapping for times pull requests were created by the user
func GetPullsCreatedTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
	pullsList []*github.PullRequest, mPulls map[int]int) {
	var time string
	for _, pull := range pullsList {
		if pull.GetUser().GetLogin() == username {
			time = pull.GetCreatedAt().Format("2006-01-02")
			util.AddToMap(mPulls, time)
		}
	}
}

// given a context, a client, a repo's owner, a repo's name, a user's login, a list of pull requests
// for a repo, and a record of pull requests for a user, adds a mapping for times pull requests
// were assigned to or mentioned user
func GetPullsEventTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
	pullsList []*github.PullRequest, mPulls map[int]int) {
	var time string
	for _, issue := range pullsList {
		num := issue.GetNumber()
		events, _, err := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
		util.ThrowError(err)
		for _, event := range events {
			if *event.Event == "assigned" && event.Assignee.GetLogin() == username ||
				*event.Event == "mentioned" && event.Actor.GetLogin() == username {
				time = event.GetCreatedAt().Format("2006-01-02")
				util.AddToMap(mPulls, time)
			}
		}
	}
}

// given a context, a client, a repo's owner, a repo's name, a user's login, a list of pull requests for a repo,
// and a record of pull requests for a user, adds a mapping for the times pull requests reviews
// were assigned to user
func GetPullsReviewRequestTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
	pullsList []*github.PullRequest, mPulls map[int]int) {
	var time string
	for _, pull := range pullsList {
		num := pull.GetNumber()
		events, _, err := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
		util.ThrowError(err)
		for _, event := range events {
			if *event.Event == "review_requested" {
				time = event.GetCreatedAt().Format("2006-01-02")
			}
		}
		reviewers, _, _ := client.PullRequests.ListReviewers(ctx, repoOwner, repoName, num, nil)
		users := reviewers.Users
		for _, reviewer := range users {
			if reviewer.GetLogin() == username {
				util.AddToMap(mPulls, time)
			}
		}
	}
}
