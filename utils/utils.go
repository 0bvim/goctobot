package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
func Colorize(color string, text string) string {
	return color + text + Reset
}

func PrintInstructions() {
	fmt.Println(Colorize(Red, "Error: 'personal_github_token' environment variable not set."))
	fmt.Println(Colorize(Magenta, "To resolve this: "))
	fmt.Println(Colorize(Green, `
      1. Generate a GitHub personal access token with the 'user:follow' and 'read:user' scopes at https://github.com/settings/tokens.
      2. Set the token in your environment with:
      export personal_github_token="your_token_here"
      3. To make this change permanent, add it to your '~/.bashrc' with:
      echo 'export personal_github_token="your_token_here"' >> ~/.bashrc
      source ~/.bashrc

      After setting up the token, you can run OctoBot commands with:
      ghbot <command> [username]

      For more details, visit the GitHub repository.`))
}

func GetToken() string {
	personalGithubToken := os.Getenv("personal_github_token")
	if personalGithubToken == "" {
		PrintInstructions()
		os.Exit(1)
	}
	return personalGithubToken
}

func GetUser(token string) string {
	// creating a new request
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// set authorization header
	req.Header.Set("Authorization", "token "+token)

	//send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	//check if response is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status doe: %d", resp.StatusCode)
	}

	// parsing responde body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	var user string
	// extract the Github login
	if login, ok := result["login"].(string); ok {
		user = login
	} else {
		log.Fatal("Login not found in the response.")
	}

	return user
}

// logFollowUnfollow logs the action of following or unfollowing a user with a timestamp.
func LogFollowUnfollow(username, action string) error {
	// Open the log file in append mode, or create it if it doesn't exist
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the current time for the timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Create the log entry
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, action, username)

	// Write the log entry to the file
	_, err = file.WriteString(logEntry)
	if err != nil {
		return err
	}

	return nil
}
