package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	Red     = "\033[31m"
	Black   = "\033[1;30m"
	Green   = "\033[1;32m"
	Magenta = "\033[1;35m"
	Reset   = "\033[0m"
)

// Function to wrap text with color codes
func Colorize(color string, text string) {
	fmt.Println(color + text + Reset)
}

func PrintHelp() {
	Colorize(Magenta, "Commands:")
	Colorize(Green, "- follow <github_user>: Follow a GitHub user")
	Colorize(Green, "- unfollow: Unfollow users that not follow you back")
	Colorize(Green, "- following: List users you're following")
	Colorize(Green, "- followers: List your followers")
}

func GetUser(token string) string {
	resp, _ := LoginRequest(token)
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	var user string
	if login, ok := result["login"].(string); ok {
		user = login
	} else {
		log.Fatal("Login not found in the response.")
	}

	return user
}

// logFollowUnfollow logs the action of following or unfollowing a user with a timestamp.
func LogFollowUnfollow(username, action string) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, action, username)

	_, err = file.WriteString(logEntry)
	if err != nil {
		return err
	}

	return nil
}

func HandleRateLimit(count *int) {
	*count++
	delay := time.Duration(60*(*count)) * time.Second
	fmt.Printf("Rate limit exceeded. Waiting for %3.f seconds...\n", delay.Seconds())
	time.Sleep(delay)
}

func GetNextURL(resp *http.Response) string {
	linkHeader := resp.Header.Get("Link")
	if linkHeader == "" {
		return ""
	}

	links := strings.Split(linkHeader, ",")
	re := regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"`)

	for _, link := range links {
		matches := re.FindStringSubmatch(link)
		if len(matches) == 3 && matches[2] == "next" {
			return matches[1]
		}
	}

	return ""
}
