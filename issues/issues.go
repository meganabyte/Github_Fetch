package issues

import (
	"context"
	"util"
	"github.com/google/go-github/github"
)

// given a list of issues created by a user and a record of issues for the user,
// adds a mapping for the time created
func GetIssueCreatedTimes(issueListCreator []*github.Issue, mIssues map[int]int) {
	var time string
	for _, issue := range issueListCreator {
		time = issue.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00")
		util.AddToMap(mIssues, time)
	}
}

// given a context, a client, a repo's owner, a repo's name, a user's login, a list of issues
// for that repo, and a record of issues for the user, adds a mapping for events where
// issue was assigned to user or user was mentioned in issue
func GetIssueEventTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string,
						username string, issueList []*github.Issue, mIssues map[int]int) {
	var time string
	for _, issue := range issueList {
		num := issue.GetNumber()
		events, _, err := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
		util.ThrowError(err)
		for _, event := range events {
			if *event.Event == "assigned" && event.Assignee.GetLogin() == username ||
				*event.Event == "mentioned" && event.Actor.GetLogin() == username {
				time = event.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00")
				util.AddToMap(mIssues, time)
			}
		}
	}
}
