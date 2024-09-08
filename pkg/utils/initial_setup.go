package initial

import(
  "fmt"
  "os"
)

// To print text with colors
const (
    Red   = "\033[31m"
    Black = "\033[1;30m"
    Green = "\033[1;32m"
    Magenta = "\033[1;35m"
    Reset = "\033[0m"
)

// Function to wrap text with color codes
func Colorize(color string, text string) string {
    return color + text + Reset
}

func GetToken() string {
  personalGithubToken := os.Getenv("personal_github_toke")
  if personalGithubToken == "" {
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
    os.Exit(0)
  }
  return personalGithubToken
}
