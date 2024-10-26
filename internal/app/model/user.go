package model

import (
	_ "embed"
	"encoding/json"
	. "fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/0bvim/goctobot/utils"
)

const (
	FOLLOWING_URL = "https://api.github.com/users/%s/following?per_page=100"
	FOLLOWERS_URL = "https://api.github.com/users/%s/followers?per_page=100"
)

//go:embed userlist.json
var userList []byte

// Main user struct
type MyUser struct {
	Followers  []User
	Following  []User
	Login      string `json:"login"`
	TargetUser string
	Token      string
	UserStatus map[string]string
}

// a single user struct
type User struct {
	Login string `json:"login"`
}

func (receiver *MyUser) PrintFollowing() {
	Println("Following:", len(receiver.Following))
}

func (receiver *MyUser) PrintFollowers() {
	Println("Followers:", len(receiver.Followers))
}

func (receiver *MyUser) PrintStatus() {
	receiver.PrintFollowers()
	receiver.PrintFollowing()
}

func (receiver *MyUser) FetchAllowDenyList() {
	err := json.Unmarshal(userList, &receiver.UserStatus)
	if err != nil {
		Printf("Error to create allow and deny list")
	}

	for key, value := range receiver.UserStatus {
		receiver.UserStatus[key] = strings.ToLower(value)
	}
}

func (receiver *MyUser) FetchFollowing(count *int) {
	var url string
	if receiver.TargetUser != "" {
		url = Sprintf(FOLLOWING_URL, receiver.TargetUser)
	} else {
		url = Sprintf(FOLLOWING_URL, receiver.Login)
	}

	fetchData(url, "following", receiver, count)
}

func (receiver *MyUser) FetchFollowers(count *int) {
	url := Sprintf(FOLLOWERS_URL, receiver.Login)
	fetchData(url, "followers", receiver, count)
}

func (receiver *MyUser) Unfollow() {
	var usersToUnfollow []string
	for _, user := range receiver.Following {
		if !userInList(user, receiver.Followers) || receiver.UserStatus[user.Login] == "allow" {
			usersToUnfollow = append(usersToUnfollow, user.Login)
		}
	}

	processUsers(usersToUnfollow, "unfollow")
}

func (receiver *MyUser) Follow() {
	if receiver.TargetUser == "" {
		Print("User to fetch? ")
		_, err := Scanln(&receiver.TargetUser)
		if err != nil {
			Println(`Error fetching target user.!`)
			os.Exit(1)
		}
	}

	if receiver.Login != receiver.TargetUser {
		receiver.FetchFollowers(new(int))
	}

	var usersToFollow []string
	for _, user := range receiver.Followers {
		if receiver.UserStatus[user.Login] == "deny" {
			continue
		}
		if !userInList(user, receiver.Following) {
			usersToFollow = append(usersToFollow, user.Login)
		}
	}

	if len(usersToFollow) > 0 {
		processUsers(usersToFollow, "follow")
	}
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
			Printf("body: %v\n", body)
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

func processUsers(users []string, command string) {
	var wg sync.WaitGroup
	count := 0
	for _, user := range users {
		wg.Add(1)
		switch command {
		case "unfollow":
			go unfollowUser(user, &count, &wg)
		case "follow":
			go followUser(user, &count, &wg)
		}
		err := utils.LogFollowUnfollow(user, command)
		if err != nil {
			Println("Error loggin"+command+"action:", err)
		}
	}
	wg.Wait()
}

func userInList(user User, list []User) bool {
	for _, u := range list {
		if u.Login == user.Login {
			return true
		}
	}

	return false
}

func unfollowUser(user string, count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := Sprintf("https://api.github.com/user/following/%s", user)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		Printf("Error unfollowing user %s: %v\n", user, err)
		return
	}
	req.Header.Set("Authorization", "token "+utils.GetToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Printf("Error unfollowing user %s: %v\n", user, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing resonse body: %v\n", err)
		}
	}(resp.Body)

	switch resp.StatusCode {
	case http.StatusNoContent:
		Printf("User: %s has been unfollowed!\n", user)
	case http.StatusForbidden, http.StatusTooManyRequests:
		utils.HandleRateLimit(count)
		unfollowUser(user, count, wg)
	default:
		Printf("Error unfollowing %s: %v\n", user, resp.Status)
	}
}

func followUser(user string, count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := Sprintf("https://api.github.com/user/following/%s", user)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		Printf("Error following user %s: %v\n", user, err)
		return
	}
	req.Header.Set("Authorization", "token "+utils.GetToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Printf("Error following user %s: %v\n", user, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing resonse body: %v\n", err)
		}
	}(resp.Body)

	switch resp.StatusCode {
	case http.StatusNoContent:
		Printf("User: %s has been followed!\n", user)
	case http.StatusForbidden, http.StatusTooManyRequests:
		utils.HandleRateLimit(count)
		followUser(user, count, wg)
	default:
		Printf("Error following %s: %v\n", user, resp.Status)
	}
}
