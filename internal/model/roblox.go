package model

type RobloxGameVotesAPI struct {
	Data []struct {
		Id        int `json:"id"`
		Upvotes   int `json:"upVotes"`
		Downvotes int `json:"downVotes"`
	} `json:"data"`
}
