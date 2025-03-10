package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
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

func GetContributors(repoPath string) ([]Contributors, error) {
	request_url := fmt.Sprintf("https://api.github.com/repos/%s/contributors", repoPath)
	github_token, _ := GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	response, err := client.R().Get(request_url)
	if err != nil {
		fmt.Println("Error getting contributors: ", err)
		return nil, err
	}

	var contributors []Contributors
	err = json.Unmarshal(response.Body(), &contributors)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err)
		return nil, err
	}
	return contributors, nil
}

func PrintRepoDetails(repo Repository, repoPath string) {
	fmt.Printf("   🔹 Name: \033[1;34m%s\033[0m\n", repo.Name)
    fmt.Printf("   📃 Description: \033[1;34m%s\033[0m\n", repo.Description)
    fmt.Printf("   💻 Main Language: %s\n", repo.Language)
    fmt.Printf("   🔒 Private: %t\n", repo.Private)

	contributors, err := GetContributors(repoPath)
	if err != nil {
		fmt.Println("Error getting contributors: ", err)
        return
	} else {
        fmt.Println("   🔒 Contributors:")
        for _, contributor := range contributors {
            contributionText := "contributions"
            if contributor.Contributions == 1 {
                contributionText = "contribution"
            }
            fmt.Printf("     🔹 \033[1;32m%s\033[0m: %d %s\n", contributor.Login, contributor.Contributions, contributionText)
        }
    }
    fmt.Printf("   🔗 URL: \033[1;34m%s\033[0m\n", repo.HTMLURL)
    fmt.Printf("   🕸️  Remote Origin: \033[1;34m%s\033[0m\n\n", repo.CloneURL)
}