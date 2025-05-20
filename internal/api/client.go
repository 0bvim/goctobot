package api

import (
	"context"
	"os"

	"github.com/google/go-github/v72/github"
)

type GitHubClient struct {
	client *github.Client
}

func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{
		client: github.NewClient(nil).WithAuthToken(os.Getenv("PERSONAL_GITHUB_TOKEN")),
	}
}

func (c *GitHubClient) GetFollowers(ctx context.Context, username string) (followers []*github.User, err error) {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	opt := &github.ListOptions{PerPage: 100}
	for {
		users, resp, err := c.client.Users.ListFollowers(ctx, *user.Login, opt)
		if err != nil {
			return nil, err
		}
		followers = append(followers, users...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return followers, nil
}

func (c *GitHubClient) GetFollowing(ctx context.Context, username string) (following []*github.User, err error) {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	opt := &github.ListOptions{PerPage: 100}
	for {
		users, resp, err := c.client.Users.ListFollowing(ctx, *user.Login, opt)
		if err != nil {
			return nil, err
		}
		following = append(following, users...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return following, nil
}
