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
		ctx := cmd.Context()
		user, err := client.GetUser(ctx)
		if err != nil {
			log.Fatalf("error getting user: %v", err)
		}

		fmt.Printf("followers: %v\t\n", *user.Followers)
		fmt.Printf("following: %v\t\n", *user.Following)
		// NOTE: Use this to fetch followers and following concurrently in follow and unfollow commands
		// var followers, following []*github.User
		// var wg sync.WaitGroup

		// wg.Add(2)
		// go func() {
		// 	defer wg.Done()
		// 	var err error
		// 	followers, err = client.GetFollowers(ctx, *user)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// }()

		// go func() {
		// 	defer wg.Done()
		// 	var err error
		// 	following, err = client.GetFollowing(ctx, *user)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// }()

		// wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
