package main

import (
	"fmt"

	"github.com/0bvim/goctobot/utils"
)

func main() {
	token := utils.GetToken() // package name to call the functions
	user := utils.GetUser(token)
	fmt.Println(user)
	fmt.Println("Success " + token)
}
