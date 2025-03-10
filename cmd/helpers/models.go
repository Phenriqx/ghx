package helpers

type Repository struct {
	Name string `json:"name"`
	Private bool `json:"private"`
	Language string `json:"language"`
	HTMLURL string `json:"html_url"`
}

type CreateRepoRequest struct {
	Name string `json:"name"`
	Private bool `json:"private"`
	Description string `json:"description"`
}