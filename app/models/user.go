package models

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func CreateUser(attributes map[string]string) *User {
	accessToken := attributes["AccessToken"]
	user := User{0, "", accessToken}
	DbMap.Insert(&user)
	return &user
}

func FindUserByAccessToken(accessToken string) *User {
	users, err := DbMap.Select(User{}, "select * from User where AccessToken = ?", accessToken)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*User)
}

func AllUsers() []*User {
	var users []*User
	rows, err := DbMap.Select(User{}, "select * from User")
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		user := row.(*User)
		users = append(users, user)
	}
	return users
}

type User struct {
	Id          int
	Login       string
	AccessToken string
}

func (user *User) Github() *github.Client {
	transport := &oauth.Transport{
		Token: &oauth.Token{AccessToken: user.AccessToken},
	}
	client := github.NewClient(transport.Client())
	return client
}

func (user *User) Save() {
	DbMap.Update(user)
}
