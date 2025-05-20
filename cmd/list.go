/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/0bvim/goctobot/internal/api"
	"github.com/google/go-github/v72/github"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List followers and following users",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewGitHubClient(os.Getenv("GITHUB_PEROSNAL_TOKEN"))
		var followers, following []*github.User
		ctx := cmd.Context()
		user, err := client.GetLogin(ctx)
		if err != nil {
			log.Fatalf("error getting user: %v", err)
		}

		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			var err error
			followers, err = client.GetFollowers(ctx, *user)
			if err != nil {
				log.Println(err)
			}
		}()

		go func() {
			defer wg.Done()
			var err error
			following, err = client.GetFollowing(ctx, *user)
			if err != nil {
				log.Println(err)
			}
		}()

		wg.Wait()

		fmt.Printf("followers: %v\t\n", len(followers))
		fmt.Printf("following: %v\t\n", len(following))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
