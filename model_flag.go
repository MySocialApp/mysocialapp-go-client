package msa

type Flag struct {
	Text        *string `json:"text"`
	Description *string `json:"description"`
	Link        *string `json:"link"`
	LinkText    *string `json:"link_text"`
	ImageUrl    *string `json:"image_url"`
}
