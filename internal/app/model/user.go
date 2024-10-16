package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	fetchData(url, "following", u, count)
}

func (u *MyUser) FetchFollowers(count *int) {
	url := fmt.Sprintf(FOLLOWERS_URL, u.Login)
	fetchData(url, "followers", u, count)
}

func (u *MyUser) Unfollow() {
	//TODO: Implement this function
}

func (u *MyUser) Follow() {
	//TODO: Implement this function
}

func fetchData(url, action string, u *MyUser, count *int) {
	for url != "" {
		resp, err := utils.FetchRequest(url)
		if err != nil {
			log.Fatalf("Error in request %s", err)
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			utils.HandleRateLimit(count)
			continue
		}

		var users []User
		if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
			body, err := io.ReadAll(resp.Body)
			fmt.Printf("body: %v\n", body)
			log.Fatalf("Error in request %s", err)
		}

		switch action {
		case "followers":
			u.Followers = append(u.Followers, users...)
		case "following":
			u.Following = append(u.Following, users...)
		}

		url = utils.GetNextURL(resp)
	}
}
