package msa

type Status struct {
	BaseWall
	Message string `json:"message"`
}

func (f *Status) GetLikes() ([]Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Status) AddLike() (*Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Status) RemoveLike() *RestError {
	// TODO
	return nil
}

func (f *Status) ListComments() ([]Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Status) AddComment(commentPost *CommentPost) (*Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Status) Delete() *RestError {
	// TODO
	return nil
}

func (f *Status) Update() *RestError {
	// TODO
	return nil
}


