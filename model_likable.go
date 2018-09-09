package msa

type Likable interface {
	GetLikes() ([]Like, *RestError)
	AddLike() (Like, *RestError)
	RemoveLike() *RestError
}
