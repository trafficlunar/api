package model

type GitHubAPI struct {
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Language    string `json:"language"`
	Url         string `json:"html_url"`
}

type GitHubData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Language    string `json:"language"`
	Url         string `json:"url"`
}
