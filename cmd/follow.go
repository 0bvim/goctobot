package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// followCmd represents the follow command
var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Follow users that follow you but you don't follow back",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("follow called")
	},
}

func init() {
	rootCmd.AddCommand(followCmd)
}
