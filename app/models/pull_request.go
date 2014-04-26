package models

import (
	"strconv"
)

type PullRequest struct {
	Id           int
	RepositoryId int
	Number       int
	State        string
	Title        string
}

func CreatePullRequest(attributes map[string]string) *PullRequest {
	repositoryId, _ := strconv.Atoi(attributes["RepositoryId"])
	number, _ := strconv.Atoi(attributes["Number"])
	pullRequest := PullRequest{
		RepositoryId: repositoryId,
		Number:       number,
		State:        attributes["State"],
		Title:        attributes["Title"],
	}
	DbMap.Insert(&pullRequest)
	return &pullRequest
}

func PullRequestsBy() []*PullRequest {
	query := SelectQuery("PullRequest", map[string]string{})

	rows, err := DbMap.Select(PullRequest{}, query)
	if err != nil {
		panic(err)
	}

	var pullRequests []*PullRequest
	for _, row := range rows {
		pullRequests = append(pullRequests, row.(*PullRequest))
	}
	return pullRequests
}
