package msa

type UrlTag struct {
	TagEntityAbstract
	OriginalUrl          *string `json:"original_url"`
	OriginalURLToDisplay *string `json:"original_url_to_display"`
	OriginalHostUrl      *string `json:"original_host_url"`
	ShortUrl             *string `json:"short_url"`
	Title                *string `json:"title"`
	Description          *string `json:"description"`
	PreviewUrl           *string `json:"preview_url"`
	StartIndex           int     `json:"start_index"`
	EndIndex             int     `json:"end_index"`
}

func (u *UrlTag) GetText() *string {
	return u.OriginalUrl
}

func (u *UrlTag) GetTextShown() *string {
	return u.OriginalURLToDisplay
}

func (u *UrlTag) GetIndices() []int {
	return []int{u.StartIndex, u.EndIndex}
}

func (u *UrlTag) SetIndices(indices []int) {
	if indices != nil && len(indices) == 2 {
		u.Indices = indices
		u.StartIndex = indices[0]
		u.EndIndex = indices[1]
	}
}
