package issues

import (
	"util"
	"github.com/google/go-github/github"
)

//
type Issues interface {
	GetIssueCreatedTimes(l []*github.Issue, mIssues map[int]int)
}

// given issues created by a user and a record of issues for the user,
// adds a mapping for the time created
func GetIssueCreatedTimes(list []*github.Issue, mIssues map[int]int) {
	var time string
	for _, issue := range list {
		time = issue.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00")
		util.AddToMap(mIssues, time)
	}
}
