package services

import (
	"context"
	"errors"
	
	"os"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

var ErrMissingToken = errors.New("GITHUB_TOKEN not defined")

// NewGitHubClient cria o cliente autenticado com o token
func NewGitHubClient() (*github.Client, context.Context, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, nil, ErrMissingToken
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client, ctx, nil
}

// Funções que usam o cliente fornecido:

func CreateRepository(client *github.Client, ctx context.Context, name string) error {
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	return err
}

func DeleteRepository(client *github.Client, ctx context.Context, owner, repo string) error {
	_, err := client.Repositories.Delete(ctx, owner, repo)
	return err
}

func ListRepositories(client *github.Client, ctx context.Context) ([]*github.Repository, error) {
	repos, _, err := client.Repositories.List(ctx, "", nil)
	return repos, err
}

func ListOpenPullRequests(client *github.Client, ctx context.Context, owner, repo string, n int) ([]*github.PullRequest, error) {
	opts := &github.PullRequestListOptions{
		State: "open",
		ListOptions: github.ListOptions{
			PerPage: n,
		},
	}
	prs, _, err := client.PullRequests.List(ctx, owner, repo, opts)
	return prs, err
}
