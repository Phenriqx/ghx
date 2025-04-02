package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phenriqx/github-cli/cmd"	
)

func main() {
	err := godotenv.Load()
	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}
	_ = GITHUB_TOKEN
	cmd.Execute()
}
