/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0bvim/goctobot/internal/api"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List followers and following users",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewGitHubClient(os.Getenv("GITHUB_PEROSNAL_TOKEN"))
		followers, err := client.GetFollowers(cmd.Context(), "0bvim")
		if err != nil {
			log.Println(err)
		}

		following, err := client.GetFollowing(cmd.Context(), "0bvim")
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("followers: %v\t\n", len(followers))
		fmt.Printf("following: %v\t\n", len(following))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
