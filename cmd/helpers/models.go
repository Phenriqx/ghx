package helpers

type Repository struct {
	Name string `json:"name"`
	Private bool `json:"private"`
	Language string `json:"language"`
	Description string `json:"description"`
	// Contributors Contributors `json:"contributors_url"`
	HTMLURL string `json:"html_url"`
	CloneURL string `json:"clone_url"`
}

type Contributors struct {
	Login string `json:"login"`
	Contributions int `json:"contributions"`
	HTMLURL string `json:"html_url"`
}

type CreateRepoRequest struct {
	Name string `json:"name"`
	Private bool `json:"private"`
	Description string `json:"description"`
}