package msa

type TagEntityAbstract struct {
	Type      *string `json:"type"`
	Text      *string `json:"text"`
	TextShown *string `json:"text_shown"`
	Indices   []int   `json:"indices"`
}

func (t *TagEntityAbstract) GetType() *string {
	return t.Type
}

func (t *TagEntityAbstract) GetText() *string {
	return t.Text
}

func (t *TagEntityAbstract) GetTextShown() *string {
	return t.TextShown
}

func (t *TagEntityAbstract) GetIndices() []int {
	return t.Indices
}

func (t *TagEntityAbstract) SetIndices(indices []int) {
	t.Indices = indices
}

type TagEntity interface {
	GetType() *string
	GetText() *string
	GetTextShown() *string
	GetIndices() []int
}
