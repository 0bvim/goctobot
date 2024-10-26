package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/0bvim/goctobot/internal/app/model"
	"github.com/0bvim/goctobot/utils"
	"github.com/joho/godotenv"
)

func init() {
	// check args
	if len(os.Args) == 1 {
		utils.PrintHelp()
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.GetToken()
}

func main() {
	command := os.Args[1]

	user := model.MyUser{}
	user.Token = utils.GetToken()
	user.Login = utils.GetUser(user.Token)
	user.FetchFollowers(new(int))
	user.FetchFollowing(new(int))
	go user.FetchAllowDenyList()

	if len(os.Args) > 2 {
		user.TargetUser = os.Args[2]
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
		user.Unfollow()
	case "followers":
		user.PrintFollowers()
	case "following":
		user.PrintFollowing()
	case "status":
		user.PrintStatus()
	case "follow":
		user.Follow()
	default:
		fmt.Println("Invalid command")
	}
}
