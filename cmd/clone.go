/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <repo-name>",
	Short: "Clone a github repository",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		repoPath := args[0]
		cloneURL := fmt.Sprintf("https://github.com/%s.git", repoPath)

		fmt.Printf("Cloning repository: %s...\n", cloneURL)

		cmdClone := exec.Command("git", "clone", cloneURL)
		cmdClone.Stdout = os.Stdout
		cmdClone.Stderr = os.Stderr
		
		err := cmdClone.Run()
		if err != nil {
			fmt.Println("Error cloning repository: ", err)
			return
		}

		fmt.Println("Repository cloned successfully!")
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
