package main

import (
	"context"
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	githuboutput "github.com/ci-space/github-output"
	"github.com/google/go-github/v73/github"
)

type Inputs struct {
	Owner string `env:"INPUT_OWNER"`
	Repo  string `env:"INPUT_REPO"`

	GithubToken string `env:"GITHUB_TOKEN"`
}

func main() {
	inputs, err := env.ParseAs[Inputs]()
	if err != nil {
		fmt.Printf("failed to read inputs: %s\n", err.Error())
		os.Exit(1)
	}

	client := newClient(inputs)

	ctx := context.Background()

	fmt.Println("fetching last commits")

	commits, _, err := client.Repositories.ListCommits(ctx, inputs.Owner, inputs.Repo, &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 1,
		},
	})
	if err != nil {
		fmt.Printf("failed to fetch commits: %s\n", err.Error())
		os.Exit(1)
	}

	if len(commits) == 0 {
		fmt.Println("repository not contains commits")
		os.Exit(1)
	}

	fmt.Println("commit fetched")

	if err = writeCommitToOutput(commits[0]); err != nil {
		fmt.Printf("failed to write commit to output: %s\n", err.Error())
		os.Exit(1)
	}
}

func writeCommitToOutput(commit *github.RepositoryCommit) error {
	return githuboutput.WriteMap(map[string]string{
		"sha":          commit.GetSHA(),
		"message":      commit.GetCommit().GetMessage(),
		"url":          commit.GetURL(),
		"author_login": commit.GetAuthor().GetLogin(),
	})
}

func newClient(inputs Inputs) *github.Client {
	client := github.NewClient(nil)

	if inputs.GithubToken != "" {
		client = client.WithAuthToken(inputs.GithubToken)
	}

	return client
}
