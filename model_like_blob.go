package msa

type LikeBlob struct {
	Total   int    `json:"total"`
	Samples []Like `json:"samples"`
}
