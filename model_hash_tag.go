package msa

type HashTag struct {
	TagEntityAbstract
	StartIndex int `json:"start_index"`
	EndIndex   int `json:"end_index"`
}

func (h *HashTag) GetIndices() []int {
	return []int{h.StartIndex, h.EndIndex}
}

func (h *HashTag) SetIndices(indices []int) {
	if indices != nil && len(indices) == 2 {
		h.StartIndex = indices[0]
		h.EndIndex = indices[1]
		h.Indices = indices
	}
}
