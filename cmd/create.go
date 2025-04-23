package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func HandleCreateRepo(repoName, description string, private bool) (*resty.Response, helpers.CreateRepoRequest, error) {
	github_token, err := helpers.GetGithubToken()
	if err != nil || github_token == "" {
		return nil, helpers.CreateRepoRequest{}, fmt.Errorf("Error fetching Github Token: %v", err)
	}
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
		return nil, helpers.CreateRepoRequest{}, err
	}
	if response.StatusCode() >= 400 {
		fmt.Printf("Error. Status Code: %v", response.StatusCode())
		return response, helpers.CreateRepoRequest{}, fmt.Errorf("Github API Error: %v", response.String())
	}

	return response, createdRepo, nil
}

func handleCreatePostRequest(repoName, desc string, private bool) {
	response, createdRepo, err := HandleCreateRepo(repoName, desc, private)
	if err != nil {
		fmt.Println("Error handling POST Request: ", err)
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
		handleCreatePostRequest(repoName, description, private)
	},
}

func init() {
	createCmd.PersistentFlags().BoolP("private", "p", false, "Create private repository.")
	createCmd.PersistentFlags().StringP("desc", "d", "", "Add repository description.")
	rootCmd.AddCommand(createCmd)
}
