package repos

import (
	"context"
	"github.com/google/go-github/github"
)

type Repos interface {
	userRepos(ctx context.Context, client *github.Client, username string)
	subscribedRepos(ctx context.Context, client *github.Client, username string)
	addToMap(m map[string]*github.Repository, list []*github.Repository)
	GetRepoList(ctx context.Context, client *github.Client, username string)
	GetRepoInfo(repo *github.Repository)
	GetCommitRepoList(ctx context.Context, client *github.Client, repoList []*github.Repository)
}

// given a context, a client, a user's login, returns a list of all repos for that user
func userRepos(ctx context.Context, client *github.Client, username string) (map[string]*github.Repository, error) {
	var repoList []*github.Repository
	emptyMap := make(map[string]*github.Repository)
	optRepo := &github.RepositoryListOptions{Type: "all", ListOptions: github.ListOptions{PerPage: 30}}
	for {
		repos, resp, err := client.Repositories.List(ctx, username, optRepo)
		if err != nil {
			return nil, err
		}
		repoList = append(repoList, repos...)
		if resp.NextPage == 0 {
			break
		}
		optRepo.Page = resp.NextPage
	}
	return AddToMap(emptyMap, repoList), nil
}

// given a context, a client, and a user's login, returns a list of repos the user is subscribed to
func subscribedRepos(ctx context.Context, client *github.Client, username string) ([]*github.Repository, error) {
	var repoList []*github.Repository
	optRepo := &github.ListOptions{PerPage: 30}
	for {
		repos, resp, err := client.Activity.ListWatched(ctx, username, optRepo)
		if err != nil {
			return nil, err
		}
		repoList = append(repoList, repos...)
		if resp.NextPage == 0 {
			break
		}
		optRepo.Page = resp.NextPage
	}
	return repoList, nil
}

// given a map of repos and a list of repos to add, adds repos from list to map if not already contained
// in map
func AddToMap(m map[string]*github.Repository, list []*github.Repository) map[string]*github.Repository {
	if len(m) == 0 {
		for _, repo := range list {
			m[repo.GetName() + repo.GetOwner().GetLogin()] = repo
		}
	} else {
		for _, repo := range list {
			key := repo.GetName() + repo.GetOwner().GetLogin()
			if _, ok := m[key]; !ok {
				m[key] = repo
			}
		}
	}
	return m
}

// given a context, a client, and a user's login, returns a list of repos to search for contributions
func GetRepoList(ctx context.Context, client *github.Client, username string) []*github.Repository {
	var repoList []*github.Repository
	map1, _ := userRepos(ctx, client, username)
	list, _ := subscribedRepos(ctx, client, username)
	map2 := AddToMap(map1, list)
	for _, v := range map2 {
		repoList = append(repoList, v)
	}
	return repoList
}

// given a repo, returns the repo's name and owner
func GetRepoInfo(repo *github.Repository) (string, string) {
	repoName := repo.GetName()
	repoOwner := repo.GetOwner().GetLogin()
	return repoName, repoOwner
}

// given a context, a client, a list of repos, returns a list of standalone repos
func GetStandaloneRepoList(ctx context.Context, client *github.Client, repoList []*github.Repository) ([]*github.Repository, error) {
	var forksList []*github.Repository
	standaloneRepoList := repoList
	optRepo := &github.RepositoryListForksOptions{ListOptions: github.ListOptions{PerPage: 30}}
	for _, repo := range repoList {
		for {
			repoName, repoOwner := GetRepoInfo(repo)
			forks, resp, err := client.Repositories.ListForks(ctx, repoOwner, repoName, optRepo)
			if err != nil {
				return nil, err
			}
			forksList = append(forksList, forks...)
			if resp.NextPage == 0 {
				break
			}
			optRepo.Page = resp.NextPage
		}
	}
	for _, fork := range forksList {
		for i := 0; i < len(standaloneRepoList); i++ {
			if fork.GetFullName() == standaloneRepoList[i].GetFullName() {
				lastElement := (len(standaloneRepoList) - 1)
				standaloneRepoList[i] = standaloneRepoList[lastElement]
				standaloneRepoList = standaloneRepoList[:lastElement]
			}
		}
	}
	return standaloneRepoList, nil
}
