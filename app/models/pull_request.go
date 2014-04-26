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
	Url          string
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

func PullRequestsBy(attributes map[string]string) []*PullRequest {
	query := SelectQuery("PullRequest", attributes)

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

func FindPullRequestBy(attributes map[string]string) *PullRequest {
	pullRequests := PullRequestsBy(attributes)
	if len(pullRequests) == 0 {
		return nil
	}
	return pullRequests[0]
}

func FindOrCreatePullRequestBy(attributes map[string]string) *PullRequest {
	pullRequest := FindPullRequestBy(attributes)
	if pullRequest != nil {
		return pullRequest
	}
	return CreatePullRequest(attributes)
}

func (pullRequest *PullRequest) Repository() *Repository {
	return FindRepositoryBy(map[string]string{"Id": strconv.Itoa(pullRequest.RepositoryId)})
}

func (pullRequest *PullRequest) Save() {
	DbMap.Update(pullRequest)
}
