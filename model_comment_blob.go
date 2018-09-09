package msa

type CommentBlob struct {
	Total   int       `json:"total"`
	Samples []Comment `json:"samples"`
}
