package models

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"strconv"
)

type User struct {
	Id          int
	Login       string
	AvatarURL   string
	AccessToken string
}

func CreateUser(attributes map[string]string) *User {
	user := User{
		Login:       attributes["Login"],
		AccessToken: attributes["AccessToken"],
	}
	DbMap.Insert(&user)
	return &user
}

func UsersBy(attributes map[string]string) []*User {
	query := SelectQuery("User", attributes)

	rows, err := DbMap.Select(User{}, query)
	if err != nil {
		panic(err)
	}

	var users []*User
	for _, row := range rows {
		users = append(users, row.(*User))
	}
	return users
}

func FindUserBy(attributes map[string]string) *User {
	users := UsersBy(attributes)
	if len(users) == 0 {
		return nil
	}
	return users[0]
}

func FindOrCreateUserBy(attributes map[string]string) *User {
	user := FindUserBy(attributes)
	if user != nil {
		return user
	}
	return CreateUser(attributes)
}

func (user *User) Github() *github.Client {
	transport := &oauth.Transport{
		Token: &oauth.Token{AccessToken: user.AccessToken},
	}
	client := github.NewClient(transport.Client())
	return client
}

func (user *User) Repositories() []*Repository {
	return RepositoriesBy(map[string]string{
		"UserId": strconv.Itoa(user.Id),
	})
}

func (user *User) PullRequests() []*PullRequest {
	return PullRequestsBy(map[string]string{})
}

func (user *User) Save() {
	DbMap.Update(user)
}
