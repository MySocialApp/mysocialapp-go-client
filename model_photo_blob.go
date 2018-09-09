package msa

type PhotoBlob struct {
	Total   int     `json:"total"`
	Samples []Photo `json:"samples"`
}
