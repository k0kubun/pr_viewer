package controllers

import (
	"github.com/revel/revel"
)

type Users struct {
	Application
}

func (c Users) Show(login string) revel.Result {
	return c.Render()
}
