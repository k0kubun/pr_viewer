package models

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
)

func CreateUser(attributes map[string]string) *User {
	user := User{
		Login:       attributes["Login"],
		AccessToken: attributes["AccessToken"],
	}
	DbMap.Insert(&user)
	return &user
}

func FindUserBy(attributes map[string]string) *User {
	query := "select * from User"
	for key, value := range attributes {
		query = fmt.Sprintf("%s where %s = '%s'", query, key, value)
	}

	users, err := DbMap.Select(User{}, query)
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
	AvatarURL   string
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
