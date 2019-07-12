package commits

import (
	"github.com/google/go-github/github"
	"util"
)

// given a list of commits for a repo opened by a user, a response body, and a map of commits,
// adds mappings for the times each commit was opened
func GetCommitList(commitsList []*github.RepositoryCommit) []*github.RepositoryCommit {
	m := make(map[string]*github.RepositoryCommit)
	var list []*github.RepositoryCommit
	for _, commit := range commitsList {
		id := commit.GetSHA()
		if _, ok := m[id]; !ok {
			m[id] = commit
		} 
	}
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

func GetCommitTimes(list []*github.RepositoryCommit, mCommits map[int]int) {
	for _, commit := range list {
		author := commit.Commit.GetAuthor()
		time := author.GetDate().Format("2006-01-02T15:04:05Z07:00")
		util.AddToMap(mCommits, time)
	}
}
