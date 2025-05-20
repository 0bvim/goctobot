package api

import (
	"github.com/google/go-github/v72/github"
)

type Users struct {
	Followers []*github.User
	Following []*github.User
}
