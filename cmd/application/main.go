package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"github.com/0bvim/goctobot/internal/app/model"
	"github.com/0bvim/goctobot/utils"
)

func init() {
	// check args
	fmt.Printf("Size = %d\n", unsafe.Sizeof(model.MyUser{}))
	if len(os.Args) == 1 {
		utils.PrintHelp()
		os.Exit(1)
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
	//TODO: implement user.FetchList to allow and deny list

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
		fmt.Printf("You have %d followers.\n", len(user.Followers))
	case "following":
		fmt.Printf("You follow %d users.\n", len(user.Following))
	case "follow":
		user.Follow()
	default:
		fmt.Println("Invalid command")
	}
}
