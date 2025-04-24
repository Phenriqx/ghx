package helpers

type Repository struct {
	Name        string `json:"name"`
	Private     bool   `json:"private"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	SSHURL      string `json:"ssh_url"`
	CloneURL    string `json:"clone_url"`
}

type Contributors struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
	HTMLURL       string `json:"html_url"`
}

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
}

type UserActivity struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
	Payload   struct {
		Size int `json:"size"`
	} `json:"payload"`
}

type GithubUser struct {
	Login string `json:"login"`
}

