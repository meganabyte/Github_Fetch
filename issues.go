package issues

import (
	"context"
	"github.com/google/go-github/github"
	"util"
)


func GetIssueCreatedTimes(issueListCreator []*github.Issue, mIssues map[int]int) {
	var time string
	for _, issue := range issueListCreator {
		time = issue.GetCreatedAt().Format("2006-01-02")
		util.AddToMap(mIssues, time)
	}
}

func GetIssueEventTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string, 
						issueList []*github.Issue, mIssues map[int]int) {
	var time string
	for _, issue := range issueList {
		num := issue.GetNumber()
		events, _, _ := client.Issues.ListIssueEvents(ctx, repoOwner, repoName, num, nil)
		for _, event := range events {
			if *event.Event == "assigned" && event.Assignee.GetLogin() == username || *event.Event == "mentioned" && event.Actor.GetLogin() == username {
				time = event.GetCreatedAt().Format("2006-01-02")
				util.AddToMap(mIssues, time)
			}
		}
	}
}






