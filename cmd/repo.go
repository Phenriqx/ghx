package cmd

import (
	"encoding/json"
	"fmt"
	// "strings"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func HandleGetRepo(repoPath string) {
	github_token, _ := helpers.GetGithubToken()
	var repo helpers.Repository
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	response, err := client.R().Get("https://api.github.com/repos/" + repoPath)
	if err != nil {
		fmt.Println("Error getting repository: ", err)	
		return
	}
	fmt.Println("https://api.github.com/repos/" + repoPath)
	err = json.Unmarshal(response.Body(), &repo)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err)
		return
	}

	helpers.PrintRepoDetails(repo, repoPath)
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo <repo-name>",
	Short: "Get information about a repository",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		repoPath := args[0]
		HandleGetRepo(repoPath)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
