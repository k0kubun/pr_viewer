package controllers

import (
	"github.com/google/go-github/github"
	"github.com/revel/revel"
	"pr_viewer/app/models"
	"pr_viewer/app/routes"
	"strconv"
)

type Users struct {
	Application
	user *models.User
}

func (c Users) Show(login string) revel.Result {
	c.user = models.FindOrCreateUserBy(map[string]string{"Login": login})
	c.RenderArgs["user"] = c.user
	if c.RenderArgs["user"] == nil {
		return c.Redirect(routes.Application.Index())
	}
	c.RenderArgs["repos"] = c.user.Repositories()
	c.loadPullRequests()
	return c.Render()
}

func (c Users) Update(login string) revel.Result {
	if c.RenderArgs["loginUser"] == nil {
		return c.Redirect(routes.Users.Show(login))
	}
	c.loginUser = c.RenderArgs["loginUser"].(*models.User)

	c.getRepositories(login)
	c.getPullRequests(login)
	return c.Redirect(routes.Users.Show(login))
}

func (c Users) loadPullRequests() {
	pullRequests := c.user.PullRequests()
	merged := []*models.PullRequest{}
	closed := []*models.PullRequest{}
	open := []*models.PullRequest{}

	for _, pullRequest := range pullRequests {
		switch pullRequest.State {
		case "merged":
			merged = append(merged, pullRequest)
		case "closed":
			closed = append(closed, pullRequest)
		case "open":
			open = append(open, pullRequest)
		}
	}
	c.RenderArgs["merged"] = merged
	c.RenderArgs["closed"] = closed
	c.RenderArgs["open"] = open
}

func (c Users) getRepositories(login string) {
	user := models.FindUserBy(map[string]string{"Login": login})
	if user == nil {
		return
	}

	githubRepositories := c.allGithubRepositories(login)
	for _, githubRepository := range githubRepositories {
		owner := githubRepository.Owner
		url := *githubRepository.HTMLURL
		if *githubRepository.Fork == true {
			githubRepositoryWithParent, _, err := c.loginUser.Github().Repositories.Get(*owner.Login, *githubRepository.Name)
			if err != nil {
				panic(err)
			}
			owner = githubRepositoryWithParent.Parent.Owner
			url = *githubRepositoryWithParent.Parent.HTMLURL
		}

		repository := models.FindOrCreateRepositoryBy(map[string]string{
			"UserId": strconv.Itoa(user.Id),
			"Name":   *githubRepository.Name,
			"Owner":  *owner.Login,
		})
		repository.Url = url
		repository.Save()
	}
}

func (c Users) allGithubRepositories(login string) []github.Repository {
	allGithubRepositories := []github.Repository{}

	for i := 1; ; i++ {
		options := &github.RepositoryListOptions{}
		options.Page = i
		githubRepositories, _, err := c.loginUser.Github().Repositories.List(login, options)
		if err != nil {
			panic(err)
		}

		if len(githubRepositories) == 0 {
			break
		}
		allGithubRepositories = append(allGithubRepositories, githubRepositories...)
	}
	return allGithubRepositories
}

func (c Users) getPullRequests(login string) {
	user := models.FindUserBy(map[string]string{"Login": login})
	if user == nil {
		return
	}

	for _, repository := range user.Repositories() {
		githubPullRequests := c.allGithubPullRequests(repository)
		c.createPullRequests(login, repository, githubPullRequests)
	}
}

func (c Users) allGithubPullRequests(repository *models.Repository) []github.PullRequest {
	allGithubPullRequests := []github.PullRequest{}

	for i := 1; ; i++ {
		options := &github.PullRequestListOptions{State: "closed"}
		options.Page = i
		githubPullRequests, res, err := c.loginUser.Github().PullRequests.List(repository.Owner, repository.Name, options)
		if err == nil {
			if len(githubPullRequests) == 0 {
				break
			}
			allGithubPullRequests = append(allGithubPullRequests, githubPullRequests...)
		} else if res.Status != "404 Not Found" {
			panic(err)
		}
	}

	for i := 1; ; i++ {
		options := &github.PullRequestListOptions{State: "open"}
		options.Page = i
		githubPullRequests, res, err := c.loginUser.Github().PullRequests.List(repository.Owner, repository.Name, options)
		if err == nil {
			if len(githubPullRequests) == 0 {
				break
			}
			allGithubPullRequests = append(allGithubPullRequests, githubPullRequests...)
		} else if res.Status != "404 Not Found" {
			panic(err)
		}
	}

	return allGithubPullRequests
}

func (c Users) createPullRequests(login string, repository *models.Repository, githubPullRequests []github.PullRequest) {
	for _, githubPullRequest := range githubPullRequests {
		requester := githubPullRequest.User
		if requester == nil {
			continue
		}
		if login != *requester.Login {
			continue
		}

		pullRequest := models.FindOrCreatePullRequestBy(map[string]string{
			"RepositoryId": strconv.Itoa(repository.Id),
			"Number":       strconv.Itoa(*githubPullRequest.Number),
		})
		pullRequest.State = *githubPullRequest.State
		if pullRequest.State == "closed" && githubPullRequest.MergedAt != nil {
			pullRequest.State = "merged"
		}
		pullRequest.Title = *githubPullRequest.Title
		pullRequest.Url = *githubPullRequest.HTMLURL
		pullRequest.Save()
	}
}
