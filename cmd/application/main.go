package main

import (
	"fmt"

	"github.com/0bvim/goctobot/utils"
)

func main() {
	token := utils.GetToken() // package name to call the functions
	fmt.Println("Success " + token)
}
