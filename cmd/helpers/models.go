package helpers

type Repository struct {
	Name string `json:"name"`
	Private bool `json:"private"`
	HTMLURL string `json:"html_url"`
}