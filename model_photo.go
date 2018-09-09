package msa

type Photo struct {
	BaseWall
	Message     *string      `json:"message"`
	Url         *string      `json:"url"`
	SmallUrl    *string      `json:"small_url"`
	MediumUrl   *string      `json:"medium_url"`
	HighUrl     *string      `json:"high_url"`
	Target      *Base        `json:"target"`
	Visible     *bool        `json:"visible"`
	TagEntities *TagEntities `json:"tag_entities"`
}


func (f *Photo) GetLikes() ([]Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Photo) AddLike() (*Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *Photo) RemoveLike() *RestError {
	// TODO
	return nil
}

func (f *Photo) ListComments() ([]Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Photo) AddComment(commentPost *CommentPost) (*Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *Photo) Delete() *RestError {
	// TODO
	return nil
}

func (f *Photo) Update() *RestError {
	// TODO
	return nil
}