package github

import (
	"context"

	"github.com/google/go-github/v50/github"
)

type Client struct {
	GithubClt *github.Client
	Config    Config
	Env       string
}

type ClientBuilder struct {
	Client *Client
}

func NewGithubClient(cfg Config) *ClientBuilder {
	return &ClientBuilder{Client: &Client{Config: cfg}}
}

func (cb *ClientBuilder) GetGithubClient() *Client {
	ctx := context.Background()
	githubClt := github.NewTokenClient(ctx, cb.Client.Config.ApiKey)

	cb.Client.GithubClt = githubClt

	return cb.Client
}
