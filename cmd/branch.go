/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage git branches",
	Long:  `Create and delete branches.`,
}

var createBranchCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new branch",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("Error getting --name flag: %v\n", err)
			return
		}
		if err := CreateBranch(name); err != nil {
			fmt.Printf("Error creating branch: %v\n", err)
		}
	},
}

var deleteBranchCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a branch",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("Error getting --name flag: %v\n", err)
			return
		}
		if err := DeleteBranch(name); err != nil {
			fmt.Printf("Error deleting branch: %v\n", err)
		}
	},
}

func CreateBranch(name string) error {
	checkout := exec.Command("git", "checkout", "-b", name)
	output, err := checkout.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error creating new branch: %v\nOutput: %v\n", err, string(output))
	}

	fmt.Printf("Created and switched to branch %s\n", name)
	return nil
}

func DeleteBranch(name string) error {
	deleteCmd := exec.Command("git", "branch", "-d", name)
	output, err := deleteCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error deleting branch %s: %v\nOutput: %v\n", name, err, string(output))
	}

	fmt.Printf("Branch %s deleted succesfully.\n", name)
	return nil
}

func init() {
	createBranchCmd.PersistentFlags().StringP("name", "n", "", "Name of the new branch")
	deleteBranchCmd.PersistentFlags().StringP("name", "n", "", "Name of the branch to delete")

	// Attach subcommands to branch
	branchCmd.AddCommand(createBranchCmd)
	branchCmd.AddCommand(deleteBranchCmd)

	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
