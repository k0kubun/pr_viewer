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
		panic(err)
		return c.Redirect(Application.Index)
	}

	c.authorizeUserByAccessToken(token.AccessToken)
	c.updateUserAttributes()
	return c.Redirect(Application.Index)
}

func (c Sessions) Destroy() revel.Result {
	for key := range c.Session {
		delete(c.Session, key)
	}
	return c.Redirect(Application.Index)
}

func (c Sessions) authorizeUserByAccessToken(accessToken string) {
	user := &models.User{AccessToken: accessToken}
	githubUser, _, err := user.Github().Users.Get("")
	if err != nil {
		panic(err)
	}
	login := *githubUser.Login
	c.Session["Login"] = login

	c.loginUser = models.FindUserBy(map[string]string{"Login": login})
	if c.loginUser == nil {
		c.loginUser = models.CreateUser(map[string]string{
			"Login": login,
		})
	}
	if c.loginUser != nil {
		c.loginUser.AccessToken = accessToken
		c.loginUser.Save()
	}
	c.RenderArgs["loginUser"] = c.loginUser
}

func (c Sessions) updateUserAttributes() {
	if c.RenderArgs["loginUser"] != nil {
		c.loginUser = c.RenderArgs["loginUser"].(*models.User)
	}
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
