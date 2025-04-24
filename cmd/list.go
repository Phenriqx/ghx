package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"

	"github.com/phenriqx/github-cli/cmd/helpers"
)

func HandleGetRequest(username string) {
	github_token, _ := helpers.GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	request_url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	response, err := client.R().Get(request_url)
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
	Use:   "list <username>",
	Short: "List all the user repositories",
	Long: `List all the repositories from the user. How to use:
			ghx {username}`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching Repositories...")
		username := args[0]
		HandleGetRequest(username)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
