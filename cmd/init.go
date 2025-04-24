/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		private, err := cmd.Flags().GetBool("private")
		if err != nil {
			fmt.Println("Error getting the --private flag: ", err)
			return
		}
		ssh, err := cmd.Flags().GetBool("ssh")
		if err != nil {
			fmt.Println("Error getting the --ssh flag: ", err)
			return
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Error getting the --name flag: ", err)
			return
		}
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			fmt.Println("Error getting the --desc flag: ", err)
			return
		}
		HandleInitCommand(private, ssh, name, desc)
	},
}

func HandleInitCommand(private, ssh bool, name, desc string) {
	fmt.Printf("Setting up environment...\n\n")

	githubToken, err := helpers.GetGithubToken()
	if err != nil {
		fmt.Printf("Error fetching Github Token: %v\n", err)
		return
	}

	username, err := helpers.GetGithubUsername(githubToken)
	if err != nil {
		fmt.Printf("Error getting username from Github: %v\n", err)
		return
	}

	cmd := exec.Command("git", "init")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running git init: %v\nOutput: %s\n", err, string(output))
		return
	}

	response, createdRepo, err := HandleCreateRepo(name, desc, private)
	if err != nil {
		fmt.Printf("Error creating Github Repo: %v\nStatus Code: %v\n", err, response.Status())
		return
	}

	var remoteOrigin string
	if ssh {
		remoteOrigin = fmt.Sprintf("git@github.com:%s/%s.git", username, createdRepo.Name)
	} else {
		remoteOrigin = fmt.Sprintf("https://github.com/%s/%s.git", username, createdRepo.Name)
	}

	addOrigin := exec.Command("git", "remote", "add", "origin", remoteOrigin)
	originOutput, err := addOrigin.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running git remote add origin command: %v\nOutput: %v\n", err, string(originOutput))
		return
	}

	fmt.Printf("Created repository: \033[31m%s\033[0m\n", response.Status())
	fmt.Println("Repository:", createdRepo.Name)
	fmt.Printf("Remote Origin: %v\n", remoteOrigin)
}

func getDirectoryName() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory name: %v\n", err)
		return "", err
	}
	return filepath.Base(path), nil
}

func init() {
	folderName, err := getDirectoryName()
	if err != nil {
		fmt.Println("Error getting current folder name: ", err)
		return
	}

	initCmd.PersistentFlags().BoolP("private", "p", false, "Create a private repository.")
	initCmd.PersistentFlags().BoolP("ssh", "s", false, "Define SSH as Remote URL, Default is HTTPS.")
	initCmd.PersistentFlags().StringP("name", "n", folderName, "Define the repository's name.")
	initCmd.PersistentFlags().StringP("desc", "d", "", "Define repository's description.")
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
