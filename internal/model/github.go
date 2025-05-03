package model

type GitHubAPI struct {
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Language    string `json:"language"`
	Url         string `json:"html_url"`
}

type GitHubData struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Language    string `json:"language"`
	Url         string `json:"url"`
}
