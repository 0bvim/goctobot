/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unfollowCmd represents the unfollow command
var unfollowCmd = &cobra.Command{
	Use:   "unfollow",
	Short: "Unfollow users that you don't follow you back",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unfollow called")
	},
}

func init() {
	rootCmd.AddCommand(unfollowCmd)
}
