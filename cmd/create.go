package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func HandleCreateRepo(repoName string, private bool, description string) {
	github_token, _ := helpers.GetGithubToken()
	var createdRepo helpers.CreateRepoRequest
	payload := helpers.CreateRepoRequest{
		Name:        repoName,
		Private:     private,
		Description: description,
	}

	client := resty.New()
	response, err := client.R().
		SetHeader("Authorization", "token "+github_token).
		SetBody(&payload).
		SetResult(&createdRepo).
		Post("https://api.github.com/user/repos")
	if err != nil {
		fmt.Println("Error creating repository: ", err)
		return
	}

	fmt.Printf("✅ Created repository \033[31m%s\033[0m\n\n", response.Status())
	fmt.Println("Repository: ", createdRepo.Name)
	fmt.Println("Private: ", createdRepo.Private)
	fmt.Println("Description: ", createdRepo.Description)
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <repo-name>",
	Short: "Creates a new repository on Github",
	Long: `Create a new Github Repository with the following syntax: 
			github-cli create {repo-name} 
			--private {include this if you want a private repo, otherwise don't} - optional. 
			--desc "provide repo description" - optional.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		private, _ := cmd.Flags().GetBool("private")    // Get the value of the --private flag
		description, _ := cmd.Flags().GetString("desc") // Get the value of the --desc flag
		HandleCreateRepo(repoName, private, description)
	},
}

func init() {
	createCmd.PersistentFlags().BoolP("private", "p", false, "Create private repository.")
	createCmd.PersistentFlags().StringP("desc", "d", "", "Add repository description.")
	rootCmd.AddCommand(createCmd)
}
