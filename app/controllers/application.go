package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/revel/revel"
	"pr_viewer/app/models"
)

var GITHUB = &oauth.Config{
	ClientId:     "e484eb4a84a86e4f8267",
	ClientSecret: "6b096615a2bfc91d9a8e8f0808216073d760f1fb",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
	RedirectURL:  "http://localhost:9000/Application/Auth",
}

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	url := GITHUB.AuthCodeURL("state")
	return c.Render(url)
}

func (c Application) Auth(code string) revel.Result {
	transport := &oauth.Transport{Config: GITHUB}
	token, err := transport.Exchange(code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(Application.Index)
	}

	accessToken := token.AccessToken
	c.Session["accessToken"] = accessToken
	c.RenderArgs["loginUser"] = models.FindUserByAccessToken(accessToken)
	if c.loginUser() == nil {
		c.RenderArgs["loginUser"] = models.CreateUser(map[string]string{
			"AccessToken": accessToken,
		})
	}
	return c.Redirect(Application.Index)
}

func (c Application) loginUser() *models.User {
	if c.RenderArgs["loginUser"] != nil {
		return c.RenderArgs["loginUser"].(*models.User)
	}
	return nil
}

func (c Application) authorize() revel.Result {
	if accessToken, ok := c.Session["accessToken"]; ok {
		user := models.FindUserByAccessToken(accessToken)
		c.RenderArgs["loginUser"] = user
	}
	return nil
}
