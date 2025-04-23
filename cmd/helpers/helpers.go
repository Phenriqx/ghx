package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func ParseDate(s string) string {
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
		return ""
	}
	formattedDate := parsedTime.Format("02/01")
	return formattedDate
}

func GetGithubToken() (string, error) {
	err := godotenv.Load("/usr/local/bin/.env")
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
		fmt.Println("Error handling GET request: ", err)
		return nil, err
	}

	var contributors []Contributors
	err = json.Unmarshal(response.Body(), &contributors)
	if err != nil {
		return nil, err
	}
	return contributors, nil
}

func PrintRepoDetails(repo Repository, repoPath string) {
	fmt.Printf("   🔹 Name: \033[1;34m%s\033[0m\n", repo.Name)
	fmt.Printf("   📃 Description: \033[1;34m%s\033[0m\n", repo.Description)
	fmt.Printf("   💻 Main Language: %s\n", repo.Language)
	fmt.Printf("   🔒 Private: %t\n", repo.Private)
	fmt.Printf("   🔗 Github URL: \033[1;34m%s\033[0m\n", repo.HTMLURL)
	fmt.Printf("   🖇️ Remote SSH URL: \033[1;34m%s\033[0m\n", repo.SSHURL)
	fmt.Printf("   🕸️ Remote HTTPS Origin: \033[1;34m%s\033[0m\n", repo.CloneURL)

	contributors, _ := GetContributors(repoPath)
	if len(contributors) == 0 {
		fmt.Println("   💡 This project does not have any contributors.")
		return
	} else {
		fmt.Println("   💡 Contributors:")
		for _, contributor := range contributors {
			contributionText := "contributions"
			if contributor.Contributions == 1 {
				contributionText = "contribution"
			}
			fmt.Printf("     🔹 \033[1;32m%s\033[0m: %d %s\n", contributor.Login, contributor.Contributions, contributionText)
		}
	}
}

func GetGithubUsername(token string) (string, error) {
	var user GithubUser
	client := resty.New()

	response, err := client.R().
		SetHeader("Authorization", "token "+token).
		SetResult(&user).
		Get("https://api.github.com/user")
	if err != nil {
		return "", err
	}
	if response.StatusCode() >= 400 {
		return "", fmt.Errorf("Github API Error: %v", response.String())
	}

	return user.Login, nil
}
