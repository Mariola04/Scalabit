package services

import (
	"context"
	"os"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

var (
	client *github.Client
	ctx    context.Context
)

func init() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN not defined on .env")
	}

	ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}


func CreateRepository(name string) error {
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	return err
}

func DeleteRepository(owner, repo string) error {
	_, err := client.Repositories.Delete(ctx, owner, repo)
	return err
}

func ListRepositories() ([]*github.Repository, error) {
	repos, _, err := client.Repositories.List(ctx, "", nil)
	return repos, err
}

// ListOpenPullRequests returns N pull requests open on a repo
func ListOpenPullRequests(owner, repo string, n int) ([]*github.PullRequest, error) {
	opts := &github.PullRequestListOptions{
		State: "open",
		ListOptions: github.ListOptions{
			PerPage: n,
		},
	}
	prs, _, err := client.PullRequests.List(ctx, owner, repo, opts)
	return prs, err
}
