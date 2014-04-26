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
	if login, ok := c.Session["Login"]; ok {
		c.loginUser = models.FindUserBy(map[string]string{"Login": login})
		c.RenderArgs["loginUser"] = c.loginUser
	}
	return nil
}

func (c Application) setLoginUrl() revel.Result {
	c.RenderArgs["loginUrl"] = GITHUB.AuthCodeURL("")
	return nil
}
