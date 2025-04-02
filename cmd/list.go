package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"

	"github.com/phenriqx/github-cli/cmd/helpers"
)

func HandleGetRequest() {
	github_token, _ := helpers.GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	response, err := client.R().Get("https://api.github.com/user/repos")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var repos []helpers.Repository
	err = json.Unmarshal(response.Body(), &repos)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err)
		return
	}
	fmt.Println("Repositories: ")
	for i, repo := range repos {
		fmt.Printf("%d 🔹 \033[1;34m%s\033[0m\n", i+1, repo.Name)
		fmt.Printf("   💻 Main Language: %s\n", repo.Language)
		fmt.Printf("   🔒 Private: %t\n", repo.Private)
		fmt.Printf("   🔗 %s\n\n", repo.HTMLURL)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the user repositories",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching Repositories...")
		HandleGetRequest()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
