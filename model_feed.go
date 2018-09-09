package msa

type Feed struct {
	Base
	Action        ActivityType   `json:"action"`
	Object        *BaseWall      `json:"object"`
	Target        *BaseWall      `json:"target"`
	AccessControl *AccessControl `json:"access_control"`
	Summary       *string        `json:"summary"`
	Owner         *User          `json:"owner"`
	Location      *Location      `json:"location"`
}

func (f *Feed) GetLikes() ([]Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Feed) AddLike() (*Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Feed) RemoveLike() *RestError {
	// TODO
	return nil
}

func (f *Feed) ListComments() ([]Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Feed) AddComment(commentPost *CommentPost) (*Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Feed) Ignore() *RestError {
	// TODO
	return nil
}

func (f *Feed) Abuse() *RestError {
	// TODO
	return nil
}

func (f *Feed) Delete() *RestError {
	// TODO
	return nil
}

func (f *Feed) Update() *RestError {
	if f.Object != nil {
		return f.Object.Update()
	}
	return nil
}

