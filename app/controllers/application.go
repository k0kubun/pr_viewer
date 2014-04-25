package controllers

import (
	"github.com/revel/revel"
	"pr_viewer/app/models"
	"pr_viewer/app/routes"
)

type Application struct {
	*revel.Controller
	loginUser *models.User
}

func (c Application) Index() revel.Result {
	if c.RenderArgs["loginUser"] != nil {
		c.loginUser = c.RenderArgs["loginUser"].(*models.User)
	}
	if c.loginUser != nil {
		return c.Redirect(routes.Users.Show(c.loginUser.Login))
	}
	return c.Render()
}

func (c Application) authorize() revel.Result {
	if accessToken, ok := c.Session["accessToken"]; ok {
		c.loginUser = models.FindUserBy(map[string]string{"AccessToken": accessToken})
		c.RenderArgs["loginUser"] = c.loginUser
	}
	return nil
}

func (c Application) setLoginUrl() revel.Result {
	c.RenderArgs["loginUrl"] = GITHUB.AuthCodeURL("")
	return nil
}

func (c Application) setUserAttributes() {
	if c.loginUser == nil {
		return
	}

	client := c.loginUser.Github()
	if client != nil {
		githubUser, _, err := client.Users.Get("")
		if err != nil {
			panic(err)
		}
		c.loginUser.Login = *githubUser.Login
		c.loginUser.AvatarURL = *githubUser.AvatarURL
		c.loginUser.Save()
	}
}
