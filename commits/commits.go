package commits

import (
	"util"
	"github.com/google/go-github/github"
)

// given a list of commits for a repo opened by a user, a response body, and a map of commits,
// adds mappings for the times each commit was opened
func GetCommitTimes(commitsList []*github.RepositoryCommit, resp *github.Response, mCommits map[int]int) {
	for _, commit := range commitsList {
		author := commit.Commit.GetAuthor()
		time := author.GetDate().Format("2006-01-02")
		util.AddToMap(mCommits, time)
	}
}
