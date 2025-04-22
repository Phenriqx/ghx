/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

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
	cmd := exec.Command("git", "init")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running git init command: ", err)
		return
	}
	fmt.Println(output)
	response, createdRepo, err := HandleCreateRepo(name, desc, private)
	if err != nil {
		fmt.Printf("Error creating Github Repo: %v\nStatus Code: %v\n", err, response.Status())
		return
	}
	ssh_url := fmt.Sprintf("git@github.com:user/%s.git", createdRepo.Name)
	addOrigin := exec.Command("git", "remote", "add", "origin", ssh_url)
	originOutput, err := addOrigin.Output()
	if err != nil {
		fmt.Println("Error running git remote add origin command: ", err)
		return
	}
	fmt.Println(originOutput)
}

func getDirectoryName() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting current directory name")
		os.Exit(1)
	}
	dirname := filepath.Dir(filename)
	name := filepath.Base(dirname)
	return name, nil
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
