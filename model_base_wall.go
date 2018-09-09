package msa

type BaseWall struct {
	Base

}

func (f *BaseWall) GetLikes() ([]Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *BaseWall) AddLike() (*Like, *RestError) {
	// TODO
	return nil, nil
}

func (f *BaseWall) RemoveLike() *RestError {
	// TODO
	return nil
}

func (f *BaseWall) ListComments() ([]Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *BaseWall) AddComment(commentPost *CommentPost) (*Comment, *RestError) {
	// TODO
	return nil, nil
}

func (f *BaseWall) Ignore() *RestError {
	// TODO
	return nil
}

func (f *BaseWall) Abuse() *RestError {
	// TODO
	return nil
}

func (f *BaseWall) Delete() *RestError {
	// TODO
	return nil
}

func (f *BaseWall) Update() *RestError {
	// TODO
	return nil
}