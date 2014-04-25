package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/revel/revel"
	"pr_viewer/app/models"
)

type Sessions struct {
	Application
}

var GITHUB = &oauth.Config{
	ClientId:     "e484eb4a84a86e4f8267",
	ClientSecret: "6b096615a2bfc91d9a8e8f0808216073d760f1fb",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
	RedirectURL:  "http://localhost:9000/auth",
}

func (c Sessions) Create(code string) revel.Result {
	transport := &oauth.Transport{Config: GITHUB}
	token, err := transport.Exchange(code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(Application.Index)
	}

	accessToken := token.AccessToken
	c.Session["accessToken"] = accessToken
	c.loginUser = models.FindUserBy(map[string]string{"AccessToken": accessToken})
	if c.loginUser == nil {
		c.loginUser = models.CreateUser(map[string]string{
			"AccessToken": accessToken,
		})
	}
	c.RenderArgs["loginUser"] = c.loginUser
	c.setUserAttributes()
	return c.Redirect(Application.Index)
}

func (c Sessions) Destroy() revel.Result {
	for key := range c.Session {
		delete(c.Session, key)
	}
	return c.Redirect(Application.Index)
}
