package commits

import (
	"github.com/google/go-github/github"
	"util"
)

func GetCommitTimes(commitsList []*github.RepositoryCommit, resp *github.Response, mCommits map[int]int) {
	for _, commit := range commitsList {
		author := commit.Commit.GetAuthor()
		time := author.GetDate().Format("2006-01-02")
		util.AddToMap(mCommits, time)
	}
}

