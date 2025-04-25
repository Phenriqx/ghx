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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := CreateBranch(name); err != nil {
			fmt.Printf("Error creating branch: %v\n", err)
		}
	},
}

var deleteBranchCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := DeleteBranch(name); err != nil {
			fmt.Printf("Error deleting branch: %v\n", err)
			return
		}
	},
}

var switchBranchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch branches",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := SwitchBranch(name); err != nil {
			fmt.Printf("Error switching to branch %s.\n", name)
			return
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

func SwitchBranch(name string) error {
	switchCmd := exec.Command("git", "checkout", name)
	output, err := switchCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error switching to branch %s.\nOutput: %v\n", name, string(output))
	}

	fmt.Printf("Switched to branch %s.\n", name)
	return nil
}

func init() {
	// Attach subcommands to branch
	branchCmd.AddCommand(createBranchCmd)
	branchCmd.AddCommand(deleteBranchCmd)
	branchCmd.AddCommand(switchBranchCmd)

	rootCmd.AddCommand(branchCmd)
}
