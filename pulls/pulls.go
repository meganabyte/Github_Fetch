package pulls

import (
	"util"
	"context"
	"github.com/google/go-github/github"
)

type Pulls interface {
	GetPullsReviewRequestTimes(ctx context.Context, list []*github.PullRequest, mPulls map[int]int)
}

// given a context, a client, a repo's owner, a repo's name, a user's login, pull requests for a repo,
// and a record of pull requests for a user, adds a mapping for times the user reviewed a pull request
func GetPullsReviewRequestTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
								list []*github.PullRequest, mPulls map[int]int) {
	var time string
	for _, pull := range list {
		num := pull.GetNumber()
		reviews, _, err := client.PullRequests.ListReviews(ctx, repoOwner, repoName, num, nil)
		if err != nil {
			return 
		}
		for _, review := range reviews {
			if review.GetUser().GetLogin() == username {
				time = review.GetSubmittedAt().Format("2006-01-02T15:04:05Z07:00")
				util.AddToMap(mPulls, time)
			}
		}
	}
}
