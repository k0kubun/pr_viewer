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
}

func (c Users) Show(login string) revel.Result {
	c.loginUser = models.FindOrCreateUserBy(map[string]string{"Login": login})
	c.RenderArgs["user"] = c.loginUser
	if c.RenderArgs["user"] == nil {
		return c.Redirect(routes.Application.Index())
	}
	c.RenderArgs["repos"] = c.loginUser.Repositories()
	c.RenderArgs["pullRequests"] = c.loginUser.PullRequests()
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

func (c Users) getRepositories(login string) {
	user := models.FindUserBy(map[string]string{"Login": login})
	if user == nil {
		return
	}

	githubRepositories, _, err := c.loginUser.Github().Repositories.List(login, nil)
	if err != nil {
		panic(err)
	}
	for _, githubRepository := range githubRepositories {
		owner := githubRepository.Owner
		if *githubRepository.Fork == true {
			githubRepositoryWithParent, _, err := c.loginUser.Github().Repositories.Get(*owner.Login, *githubRepository.Name)
			if err != nil {
				panic(err)
			}
			owner = githubRepositoryWithParent.Parent.Owner
		}
		models.FindOrCreateRepositoryBy(map[string]string{
			"UserId": strconv.Itoa(user.Id),
			"Name":   *githubRepository.Name,
			"Owner":  *owner.Login,
		})
	}
}

func (c Users) getPullRequests(login string) {
	user := models.FindUserBy(map[string]string{"Login": login})
	if user == nil {
		return
	}

	for _, repository := range user.Repositories() {
		options := &github.PullRequestListOptions{State: "closed"}
		githubPullRequests, _, err := c.loginUser.Github().PullRequests.List(repository.Owner, repository.Name, options)
		if err != nil {
			panic(err)
		}
		for _, githubPullRequest := range githubPullRequests {
			models.FindOrCreatePullRequestBy(map[string]string{
				"RepositoryId": strconv.Itoa(repository.Id),
				"Number":       strconv.Itoa(*githubPullRequest.Number),
			})
		}
	}
}
