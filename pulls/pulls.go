package pulls

import (
	"util"
	"context"
	"github.com/google/go-github/github"
)

// given a context, a client, a repo's owner, a repo's name, a user's login, a list of pull requests for a repo,
// and a record of pull requests for a user, adds a mapping for the times pull requests reviews
// were assigned to user
func GetPullsReviewRequestTimes(ctx context.Context, client *github.Client, repoOwner string, repoName string, username string,
								pullsList []*github.PullRequest, mPulls map[int]int) {
	 var time string
	 for _, pull := range pullsList {
		num := pull.GetNumber()
		reviews, _, err := client.PullRequests.ListReviews(ctx, repoOwner, repoName, num, nil)
		util.ThrowError(err)
		for _, review := range reviews {
			if review.GetUser().GetLogin() == username {
				time = review.GetSubmittedAt().Format("2006-01-02T15:04:05Z07:00")
				util.AddToMap(mPulls, time)
			}
		}
	}	
}
