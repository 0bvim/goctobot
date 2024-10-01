// package main
//
// import (
// 	"fmt"
//
// 	"github.com/0bvim/goctobot/utils"
// )
//
// func main() {
// 	token := utils.GetToken() // package name to call the functions
// 	user := utils.GetUser(token)
// 	fmt.Println(user)
// 	fmt.Println("Success " + token)
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/0bvim/goctobot/utils"
)

var personalGithubToken string

type User struct {
	Login string `json:"login"`
}

func init() {
	// Check if the token is set
	personalGithubToken = os.Getenv("personal_github_token")
	if personalGithubToken == "" {
		fmt.Println(`Error: 'personal_github_token' environment variable not set.
To resolve this:
1. Generate a GitHub personal access token with the 'user:follow' and 'read:user' scopes.
2. Set the token in your environment with:
   export personal_github_token="your_token_here"
3. To make this change permanent, add it to your '~/.bashrc' with:
   echo 'export personal_github_token="your_token_here"' >> ~/.bashrc
   source ~/.bashrc

After setting up the token, you can run OctoBot commands with:
   ghbot <command> [username]
`)
		os.Exit(1)
	}
}

func handleRateLimit(count *int) {
	*count++
	delay := time.Duration(60*(*count)) * time.Second
	fmt.Printf("Rate limit exceeded. Waiting for %03d seconds...\n", delay)
	time.Sleep(delay)
}

func fetchData(url string, count *int) ([]User, error) {
	var allData []User

	for url != "" {
		resp, err := makeRequest(url)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			handleRateLimit(count)
			continue
		}

		var users []User
		if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
			return nil, err
		}

		allData = append(allData, users...)
		url = getNextURL(resp)
	}
	return allData, nil
}

func fetchFollowers(user string, count *int) ([]User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/followers?per_page=100", user)
	return fetchData(url, count)
}

func fetchFollowing(user string, count *int) ([]User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/following?per_page=100", user)
	return fetchData(url, count)
}

func followUser(user string, count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.github.com/user/following/%s", user)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Printf("Error following user %s: %v\n", user, err)
		return
	}
	req.Header.Set("Authorization", "token "+personalGithubToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error following user %s: %v\n", user, err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNoContent:
		fmt.Printf("User: %s has been followed!\n", user)
	case http.StatusForbidden, http.StatusTooManyRequests:
		handleRateLimit(count)
		followUser(user, count, wg)
	default:
		fmt.Printf("Error following %s: %v\n", user, resp.Status)
	}
}

func unfollowUser(user string, count *int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.github.com/user/following/%s", user)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("Error unfollowing user %s: %v\n", user, err)
		return
	}
	req.Header.Set("Authorization", "token "+personalGithubToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error unfollowing user %s: %v\n", user, err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNoContent:
		fmt.Printf("User: %s has been unfollowed!\n", user)
	case http.StatusForbidden, http.StatusTooManyRequests:
		handleRateLimit(count)
		unfollowUser(user, count, wg)
	default:
		fmt.Printf("Error unfollowing %s: %v\n", user, resp.Status)
	}
}

func processUsers(users []string, command string) {
	var wg sync.WaitGroup
	count := 0
	for _, user := range users {
		wg.Add(1)
		if command == "unfollow" {
			go unfollowUser(user, &count, &wg)
		} else if command == "follow" {
			go followUser(user, &count, &wg)
		}
		err := utils.LogFollowUnfollow(user, command)
		if err != nil {
			fmt.Println("Error loggin"+command+"action:", err)
		}
	}
	wg.Wait()
}

func makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+personalGithubToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getNextURL(resp *http.Response) string {
	linkHeader := resp.Header.Get("Link")
	if linkHeader == "" {
		return ""
	}

	// Split the header by comma to handle multiple links
	links := strings.Split(linkHeader, ",")

	// Regular expression to capture the URL and its rel type
	re := regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"`)

	// Iterate over each link
	for _, link := range links {
		matches := re.FindStringSubmatch(link)
		if len(matches) == 3 && matches[2] == "next" {
			// Return the URL if the rel type is "next"
			return matches[1]
		}
	}

	// Return an empty string if no "next" link is found
	return ""
}

func main() {
	command := os.Args[1]
	var targetUser string
	if len(os.Args) > 2 {
		targetUser = os.Args[2]
	}

	// Capture termination signals
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Cleaning up...")
		os.Exit(1)
	}()

	switch command {
	case "unfollow":
		following, _ := fetchFollowing(utils.GetUser(personalGithubToken), new(int))
		followers, _ := fetchFollowers(utils.GetUser(personalGithubToken), new(int))
		var usersToUnfollow []string
		for _, user := range following {
			if !userInList(user, followers) {
				usersToUnfollow = append(usersToUnfollow, user.Login)
			}
		}
		processUsers(usersToUnfollow, "unfollow")
	case "followers":
		followers, _ := fetchFollowers(utils.GetUser(personalGithubToken), new(int))
		fmt.Printf("You have %d followers.\n", len(followers))
	case "following":
		following, _ := fetchFollowing(utils.GetUser(personalGithubToken), new(int))
		fmt.Printf("You follow %d users.\n", len(following))
	case "follow":
		if targetUser == "" {
			fmt.Print("User to fetch? ")
			fmt.Scanln(&targetUser)
		}
		followers, _ := fetchFollowers(targetUser, new(int))
		following, _ := fetchFollowing(utils.GetUser(personalGithubToken), new(int))
		var usersToFollow []string
		for _, user := range followers {
			if !userInList(user, following) {
				usersToFollow = append(usersToFollow, user.Login)
			}
		}
		processUsers(usersToFollow, "follow")
	default:
		fmt.Println("Invalid command")
	}
}

func userInList(user User, list []User) bool {
	for _, u := range list {
		if u.Login == user.Login {
			return true
		}
	}
	return false
}
