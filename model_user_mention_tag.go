package msa

type UserMentionTag struct {
	TagEntityAbstract
	MentionedUser *User `json:"mentioned_user"`
	Target        *Feed `json:"target"`
	StartIndex    int   `json:"start_index"`
	EndIndex      int   `json:"end_index"`
}

func (u *UserMentionTag) GetText() *string {
	if u.MentionedUser == nil {
		return nil
	}
	return u.MentionedUser.DisplayedName
}

func (u *UserMentionTag) GetTextShown() *string {
	// TODO
	return nil
}

func (u *UserMentionTag) GetIndices() []int {
	return []int{u.StartIndex, u.EndIndex}
}

func (u *UserMentionTag) SetIndices(indices []int) {
	if indices != nil && len(indices) == 2 {
		u.Indices = indices
		u.StartIndex = indices[0]
		u.EndIndex = indices[1]
	}
}
