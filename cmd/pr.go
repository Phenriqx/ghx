/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v50/github"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Manage Pull Requests",
}

var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all open PRs",
	Run: func(cmd *cobra.Command, args []string) {
		githubToken, err := helpers.GetGithubToken()
		if err != nil {
			fmt.Printf("Error getting Github Token: %v\n", err)
			return
		}
		githubClient := getGithubClient(githubToken)
		owner, repoName, err := helpers.GetRepoInfo()
		if err != nil {
			fmt.Printf("Error getting repository's information: %v\n", err)
			return
		}

		prs, response, err := githubClient.PullRequests.List(context.Background(), owner, repoName, nil)
		if err != nil {
			fmt.Printf("Error getting PRs from Github: %v, %v\n", err, response.Status)
			return
		}
		if len(prs) == 0 {
			fmt.Printf("No open Pull Requests on this repository.")
			return
		}

		for _, pr := range prs {
			fmt.Printf("#%d %s (%s)\n", *pr.Number, *pr.Title, *pr.State)
			fmt.Printf("Is meargeable: %v\n", *pr.Mergeable)
			if pr.GetMerged() {
				fmt.Printf("Merged By %v at %v\n", *pr.MergedBy, *pr.MergedAt)
			} else {
				fmt.Printf("PR not merged yet.")
			}
		}
	},
}

var (
	prBody string
	prHead string
	prBase string
)

var prCreateCmd = &cobra.Command{
	Use:   "new",
	Short: "Open a new Pull Requests",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prTitle := args[0]
		if err := createPullRequest(prTitle); err != nil {
			fmt.Printf("Error creating pull request: %v\n", err)
			return
		}
	},
}

var prMessage string

var prMergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge a Pull Request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Missing PR number argument.\n")
			return
		}

		prNumber := args[0]
		parsed, err := strconv.Atoi(prNumber)
		if err != nil {
			fmt.Printf("Error parsing PR number: %v\nPlease enter a valid number!", err)
			return
		}
		if parsed == 0 {
			fmt.Printf("Invalid PR number. Please enter a valid number.")
			return
		}

		if err := mergePullRequest(parsed); err != nil {
			fmt.Printf("Error deleting pull request: %v\n", err)
			return
		}
	},
}

func createPullRequest(prTitle string) error {
	githubToken, err := helpers.GetGithubToken()
	if err != nil {
		return fmt.Errorf("Error getting Github Token: %v\n", err)
	}
	client := getGithubClient(githubToken)
	owner, repoName, err := helpers.GetRepoInfo()
	if err != nil {
		return fmt.Errorf("Error getting repository's information: %v\n", err)
	}

	newPr := &github.NewPullRequest{
		Title: github.String(prTitle),
		Body:  github.String(prBody),
		Head:  github.String(prHead),
		Base:  github.String(prBase),
	}
	pr, response, err := client.PullRequests.Create(context.Background(), owner, repoName, newPr)
	if err != nil {
		return fmt.Errorf("Error creating new pull request: %v\nStatus: %v\n", err, response.Status)
	}

	fmt.Printf("New Pull Request created!\n")
	fmt.Printf("URL: %s\n", pr.GetHTMLURL())
	fmt.Printf("ID: %d\n", pr.GetID())

	return nil
}

func mergePullRequest(prNumber int) error {
	token, err := helpers.GetGithubToken()
	if err != nil {
		return fmt.Errorf("Error getting Github Token: %v\n", err)
	}

	client := getGithubClient(token)
	owner, repoName, err := helpers.GetRepoInfo()
	if err != nil {
		return fmt.Errorf("Error getting repository's information: %v\n", err)
	}

	mergeResult, response, err := client.PullRequests.Merge(
		context.Background(),
		owner,
		repoName,
		prNumber,
		prMessage,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Error merging PR with number %d: %v\nStatus Code: %v\n", prNumber, err, response.Status)
	}

	if mergeResult.GetMerged() {
		fmt.Printf("Pull Request %d merged successfully\n", prNumber)
	} else {
		fmt.Printf("Could not merge Pull Request %d: %v\n", prNumber, mergeResult.GetMessage())
	}

	return nil
}

func getGithubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	return githubClient
}

func init() {
	prCreateCmd.PersistentFlags().StringVar(&prBody, "body", "", "Description of the pull request")
	prCreateCmd.PersistentFlags().StringVar(&prHead, "head", "", "Branch with your changes (required)")
	prCreateCmd.PersistentFlags().StringVar(&prBase, "base", "", "	Branch you want to merge into")

	prMergeCmd.PersistentFlags().StringVar(&prMessage, "message", "", "Enter a message for your PR.")

	prCreateCmd.MarkFlagRequired("head")
	prCmd.AddCommand(prCreateCmd)
	rootCmd.AddCommand(prCmd)
}
