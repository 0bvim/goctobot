package api

import (
	"context"
	"fmt"

	"github.com/google/go-github/v72/github"
)

func (c *GitHubClient) GetLogin(ctx context.Context) (*string, error) {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return user.Login, nil
}

func (c *GitHubClient) GetUser(ctx context.Context) (*github.User, error) {
	user, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *GitHubClient) GetFollowers(ctx context.Context, username string) ([]*github.User, error) {
	user, err := c.GetLogin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated user: %w", err)
	}

	opt := &github.ListOptions{PerPage: 100}
	_, resp, err := c.client.Users.ListFollowers(ctx, *user, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to get first page: %w", err)
	}

	totalPages := resp.LastPage
	if totalPages == 0 {
		return []*github.User{}, nil
	}

	sem := make(chan struct{}, 10)
	var cp ConcurrentProcessing
	for page := 1; page <= totalPages; page++ {
		cp.wg.Add(1)
		go func(page int) {
			defer cp.wg.Done()
			sem <- struct{}{}        // concurrency limiter "Acquire"
			defer func() { <-sem }() // concurrency release

			pageOpt := &github.ListOptions{
				Page:    page,
				PerPage: opt.PerPage,
			}

			users, _, err := c.client.Users.ListFollowers(ctx, *user, pageOpt)
			if err != nil {
				cp.mu.Lock()
				cp.errors = append(cp.errors, fmt.Errorf("page %v: %w", page, err))
				cp.mu.Unlock()
			}

			cp.mu.Lock()
			cp.allUsers = append(cp.allUsers, users...)
			cp.mu.Unlock()
		}(page)
	}

	cp.wg.Wait()
	if len(cp.errors) > 0 {
		return nil, fmt.Errorf("errors occurred while fetching users: %v", cp.errors)
	}

	return cp.allUsers, nil
}

func (c *GitHubClient) GetFollowing(ctx context.Context, username string) ([]*github.User, error) {
	user, err := c.GetLogin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated user: %w", err)
	}

	opt := &github.ListOptions{PerPage: 100}
	_, resp, err := c.client.Users.ListFollowing(ctx, *user, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to get first page: %w", err)
	}

	totalPages := resp.LastPage
	if totalPages == 0 {
		return []*github.User{}, nil
	}

	sem := make(chan struct{}, 10)
	var cp ConcurrentProcessing
	for page := 1; page <= totalPages; page++ {
		cp.wg.Add(1)
		go func(page int) {
			defer cp.wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			pageOpt := &github.ListOptions{
				Page:    page,
				PerPage: opt.PerPage,
			}

			users, _, err := c.client.Users.ListFollowing(ctx, *user, pageOpt)
			if err != nil {
				cp.mu.Lock()
				cp.errors = append(cp.errors, fmt.Errorf("page %v: %w", page, err))
				cp.mu.Unlock()
			}

			cp.mu.Lock()
			cp.allUsers = append(cp.allUsers, users...)
			cp.mu.Unlock()
		}(page)
	}

	cp.wg.Wait()
	if len(cp.errors) > 0 {
		return nil, fmt.Errorf("errors occurred while fetching users: %v", cp.errors)
	}

	return cp.allUsers, nil
}
