package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/revel/revel"
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
	transport = &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}

	client := github.NewClient(transport.Client())
	repos, _, _ := client.Repositories.List("octocat", nil)
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
	return c.Redirect(Application.Index)
}
