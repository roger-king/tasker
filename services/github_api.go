package services

import (
	"context"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type GithubAPIService struct {
	Client *github.Client
}

func NewGithubAPIService(accessToken string) *GithubAPIService {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubAPIService{
		Client: client,
	}
}

func (g *GithubAPIService) GetUser() (*github.User, error) {
	ctx := context.Background()
	user, _, err := g.Client.Users.Get(ctx, "")

	if err != nil {
		return nil, err
	}

	return user, nil
}
