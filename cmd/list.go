/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
		fmt.Printf("   🔒 Private: %t\n", repo.Private)
		fmt.Printf("   🔗 %s\n\n", repo.HTMLURL)
	}
	// fmt.Println(string(response.Body()))
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching Repositories...")
		HandleGetRequest()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
