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
	Short: "Set up the Git/Github environment automatically.",
	Long: `Set up the Git/Github environment automatically for you. 
It initializes a local repository, creates a Github repository with the current folder's name if none are provided and also sets up HTTPS as the Remote Origin if you don't opt for SSH.
How to use: ghx init 
		--name "enter a name for your repo"
		--private - creates a private repo
		--ssh - defines the remote origin as SSH .
		--desc "provide a description for your repo" 
		--push - auto pushes to Github
		--readme - initializes the repo with a README.md file
		--gitignore - initializes the repo with a .gitignore file
All flags are optional.`,

	Run: func(cmd *cobra.Command, args []string) {
		private, ssh, gitignore, readme, push, name, desc, err := parseFlags(cmd)
		if err != nil {
			fmt.Printf("Error parsing flags: %v\n", err)
			return
		}

		HandleInitCommand(private, ssh, gitignore, readme, push, name, desc)
	},
}

func HandleInitCommand(private, ssh, gitignore, readme, push bool, name, desc string) {
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

	if err := optionalFiles(gitignore, readme, push, createdRepo); err != nil {
		fmt.Printf("Error creating optional file: %v\n", err)
		return
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

func optionalFiles(gitignore, readme, push bool, createdRepo helpers.CreateRepoRequest) error {
	if gitignore {
		if err := createGitignoreFile(); err != nil {
			fmt.Printf("Error creating .gitignore file: %v\n", err)
			return err
		}
	}
	if readme {
		if err := createReadmeFile(createdRepo.Name); err != nil {
			fmt.Printf("Error creating README.md file: %v\n", err)
			return err
		}
	}
	if push {
		if err := autoPushInitialCommit(createdRepo.DefaultBranch); err != nil {
			fmt.Printf("Error auto pushing repo to Github: %v\n", err)
			return err
		}
	}

	return nil
}

func getDirectoryName() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory name: %v\n", err)
		return "", err
	}
	return filepath.Base(path), nil
}

func autoPushInitialCommit(defaultBranch string) error {
	cmd := exec.Command("git", "commit", "-am", "initial commit")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error on git commit: %v\nOutput: %v\n", err, string(output))
	}
	pushCmd := exec.Command("git", "push", "-u", "origin", defaultBranch)
	pushOutput, err := pushCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error pushing code to Github: %v\nOutput: %v\n", err, string(pushOutput))
	}

	return nil
}

func createReadmeFile(repoName string) error {
	var content string = fmt.Sprintf("# %v\n\nProject initialized using ghx", repoName)
	return os.WriteFile("README.md", []byte(content), 0644)
}

func createGitignoreFile() error {
	content := `
	### BINARIES ###
	*.exe
	*.dll
	*.so	
	*.dylib
	*.test
	
	### DOTENV ###
	*.env

	### IDE / Editor configs ###
	.vscode/
	.idea/
	*.swp
	`

	return os.WriteFile(".gitignore", []byte(content), 0644)
}

func parseFlags(cmd *cobra.Command) (private, ssh, gitignore, readme, push bool, name, desc string, err error) {
	private, err = cmd.Flags().GetBool("private")
	if err != nil {
		return
	}

	ssh, err = cmd.Flags().GetBool("ssh")
	if err != nil {
		return
	}

	gitignore, err = cmd.Flags().GetBool("gitignore")
	if err != nil {
		return
	}

	readme, err = cmd.Flags().GetBool("readme")
	if err != nil {
		return
	}

	push, err = cmd.Flags().GetBool("push")
	if err != nil {
		return
	}

	name, err = cmd.Flags().GetString("name")
	if err != nil {
		return
	}

	desc, err = cmd.Flags().GetString("desc")
	if err != nil {
		return
	}

	return
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
	initCmd.PersistentFlags().Bool("gitignore", false, "Initializes your repository already with a .gitignore file.")
	initCmd.PersistentFlags().Bool("readme", false, "Initializes your repository with a README.md file.")
	initCmd.PersistentFlags().Bool("push", false, "Auto push after initializing repo.")
	rootCmd.AddCommand(initCmd)
}
