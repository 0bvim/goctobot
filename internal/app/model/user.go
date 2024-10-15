package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0bvim/goctobot/utils"
)

const (
	FOLLOWING_URL = "https://api.github.com/users/%s/following?per_page=100"
	FOLLOWERS_URL = "https://api.github.com/users/%s/followers?per_page=100"
)

// Main user struct
type MyUser struct {
	Followers  []User
	Following  []User
	Allowed    []User
	Denied     []User
	Login      string `json:"login"`
	TargetUser string
	Token      string
}

// a single user struct
type User struct {
	Login string `json:"login"`
}

func (u *MyUser) FetchFollowing(count *int) {
	url := fmt.Sprintf(FOLLOWING_URL, u.Login)
	u.fetchData(url, count)
}

func (u *MyUser) FetchFollowers(count *int) {
	url := fmt.Sprintf(FOLLOWERS_URL, u.Login)
	u.fetchData(url, count)
}

func (u *MyUser) Unfollow() {
	//TODO: Implement this function
}

func (u *MyUser) Follow() {
	//TODO: Implement this function
}

func (u *MyUser) fetchData(url, action string, count *int) ([]User, error) {
	for url != "" {
		resp, err := utils.RequestMaker(u.Token)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			utils.HandleRateLimit(count)
			continue
		}

		var users []User
		if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
			return nil, err
		}

		if action == "follow" {
			u.Followers = append(u.Followers, users...)
		} else if action == "unfollow" {
			u.Following = append(u.Following, users...)
		}

		// url getNextUrl implement
	}
	return nil, nil
}
