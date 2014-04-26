package models

import (
	"strconv"
)

type Repository struct {
	Id              int
	UserId          int
	Owner           string
	Name            string
	Url             string
	Contributed     bool
	StargazersCount int
}

func CreateRepository(attributes map[string]string) *Repository {
	userId, _ := strconv.Atoi(attributes["UserId"])
	repository := Repository{
		UserId: userId,
		Name:   attributes["Name"],
		Owner:  attributes["Owner"],
	}
	DbMap.Insert(&repository)
	return &repository
}

func RepositoriesBy(attributes map[string]string) []*Repository {
	query := SelectQuery("Repository", attributes)

	rows, err := DbMap.Select(Repository{}, query)
	if err != nil {
		panic(err)
	}

	var repositories []*Repository
	for _, row := range rows {
		repositories = append(repositories, row.(*Repository))
	}
	return repositories
}

func FindRepositoryBy(attributes map[string]string) *Repository {
	repositories := RepositoriesBy(attributes)
	if len(repositories) == 0 {
		return nil
	}
	return repositories[0]
}

func FindOrCreateRepositoryBy(attributes map[string]string) *Repository {
	repository := FindRepositoryBy(attributes)
	if repository != nil {
		return repository
	}
	return CreateRepository(attributes)
}

func (repository *Repository) Save() {
	DbMap.Update(repository)
}
