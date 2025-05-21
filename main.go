/*
Copyright © 2025 niviciusdev@gmail.com
*/
package main

import (
	"log"

	"github.com/0bvim/goctobot/cmd"
	"github.com/0bvim/goctobot/internal/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var env config.Config
	err = config.LoadConfig(&env)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	cmd.Execute()
}
