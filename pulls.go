package pulls

import (
	"context"
	"github.com/google/go-github/github"
	"util"
)

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

func GetPullsEventTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, 
			pullsList []*github.PullRequest, mPulls map[int]int) {
	var time string
	for _, issue := range pullsList {
		num := issue.GetNumber()
		events, _, _ := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
		for _, event := range events {
			if *event.Event == "assigned" && event.Assignee.GetLogin() == username || 
			   *event.Event == "mentioned" && event.Actor.GetLogin() == username {
				time = event.GetCreatedAt().Format("2006-01-02")
				util.AddToMap(mPulls, time)
			}
		}
	}
}

func GetPullsReviewRequestTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, 
				pullsList []*github.PullRequest,  mPulls map[int]int) {
	var time string
	for _, pull := range pullsList {
		num := pull.GetNumber()
		events, _, _ := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
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
