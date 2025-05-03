package model

type GitHubAPI struct {
	Description string `json:"description"`
	Stars       uint32 `json:"stargazers_count"`
	Language    string `json:"language"`
	Url         string `json:"html_url"`
}

type GitHubData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       string `json:"stars"`
	Language    string `json:"language"`
	Url         string `json:"url"`
}
