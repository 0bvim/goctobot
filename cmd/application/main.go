package main

import (
	"fmt"

	"github.com/go_rest_api/pkg/utils"
)

func main() {
	token := utils.GetToken() // package name to call the functions
	fmt.Println("Success " + token)
}
