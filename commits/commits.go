package commits

import (
	"github.com/google/go-github/github"
	"util"
)

//
type Commits interface {
	GetCommitTimes(list []*github.RepositoryCommit, mCommits map[int]int)
}

// given commits for a repository, adds mapping between a commit and the time
// the commit was created
func GetCommitTimes(list []*github.RepositoryCommit, mCommits map[int]int) {
	for _, commit := range list {
		author := commit.Commit.GetAuthor()
		time := author.GetDate().Format("2006-01-02T15:04:05Z07:00")
		util.AddToMap(mCommits, time)
	}
}
