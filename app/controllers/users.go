package controllers

import (
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
	return c.Render()
}

func (c Users) Update(login string) revel.Result {
	user := models.FindUserBy(map[string]string{"Login": login})
	if user == nil {
		return c.Redirect(routes.Users.Show(login))
	}

	githubRepositories, _, err := user.Github().Repositories.List(login, nil)
	if err != nil {
		panic(err)
	}
	for _, githubRepository := range githubRepositories {
		models.FindOrCreateRepositoryBy(map[string]string{
			"UserId": strconv.Itoa(user.Id),
			"Name":   *githubRepository.Name,
		})
	}
	return c.Redirect(routes.Users.Show(login))
}
