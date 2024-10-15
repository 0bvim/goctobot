package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func printInvalidToken() {
	fmt.Println(Colorize(Red, "Error: 'personal_github_token' environment variable not set."))
	fmt.Println(Colorize(Magenta, "To solve this: "))
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

func ValidToken(token string) error {
	resp, err := requestMaker(token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token validation failed: received status code %d", resp.StatusCode)
	}

	return nil
}

func GetToken() string {
	personalGithubToken := os.Getenv("personal_github_token")
	err := ValidToken(personalGithubToken)
	if err != nil {
		printInvalidToken()
		os.Exit(1)
	}
	return personalGithubToken
}

func requestMaker(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token validation failed: received status code %d", resp.StatusCode)
	}

	return resp, nil
}
