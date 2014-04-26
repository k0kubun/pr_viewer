package controllers

import (
	"github.com/revel/revel"
	"pr_viewer/app/models"
	"pr_viewer/app/routes"
)

type Users struct {
	Application
}

func (c Users) Show(login string) revel.Result {
	c.RenderArgs["user"] = models.FindOrCreateUserBy(map[string]string{"Login": login})
	if c.RenderArgs["user"] == nil {
		return c.Redirect(routes.Application.Index())
	}
	return c.Render()
}
