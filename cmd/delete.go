/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func HandleDeleteRepo(repoPath string) {
	github_token, _ := helpers.GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token).SetHeader("Accept", "application/vnd.github.v3+json")
	request_url := fmt.Sprintf("https://api.github.com/repos/%s", repoPath)

	response, err := client.R().Delete(request_url)
	if err != nil {
		fmt.Println("Error handling delete request", err)
		return
	}

	if response.StatusCode() == 204 {
		fmt.Println("Repository deleted successfully")
	} else {
		fmt.Println("Error deleting repository", response.StatusCode())
	}
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <repo-name>",
	Short: "Deletes a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := args[0]
		HandleDeleteRepo(repoPath)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
