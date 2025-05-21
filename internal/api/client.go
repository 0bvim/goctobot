package api

import (
	"os"
	"sync"

	"github.com/google/go-github/v72/github"
)

type GitHubClient struct {
	client *github.Client
}

type ConcurrentProcessing struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	allUsers []*github.User
	errors   []error
}

func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{
		client: github.NewClient(nil).WithAuthToken(os.Getenv("PERSONAL_GITHUB_TOKEN")),
	}
}
