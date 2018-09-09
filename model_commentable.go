package msa

type Commentable interface {
	ListComments() []Comment
	AddComment(commentPost ModelCommentPost) Comment
	GetCommentsTotal() int
	GetCommentsSamples() []Comment
}
