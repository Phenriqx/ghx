/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func HandleSearchGetRequest(search string) {
	github_token, _ := helpers.GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	response, err := client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		Get("https://api.github.com/search/repositories?q=" + search)
	if err != nil {
		fmt.Println("Error getting repositories: ", err)
		return
	}

	var results struct {
		Items []helpers.Repository `json:"items"`
	}
	err = json.Unmarshal(response.Body(), &results)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err)
		return
	}

    for i, repo := range results.Items[:5] {
		fmt.Printf("\n%d %s (%d⭐)\n", i+1, repo.Name, repo.Stars)
		fmt.Printf("   📜 %s\n", repo.Description)
		fmt.Printf("   🔗 %s\n", repo.HTMLURL)
		fmt.Printf("   🖥️ Language: %s\n", repo.Language)
	}
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <repository>",
	Short: "Search for a repository on Github",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search := args[0]
		fmt.Println("Searching for: ", search)
		HandleSearchGetRequest(search)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
