package helpers

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func GetGithubToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
        return "", fmt.Errorf("GITHUB_TOKEN environment variable not found")
    }
	return githubToken, nil
}