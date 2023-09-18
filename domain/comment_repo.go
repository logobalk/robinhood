package domain

//go:generate mockery --name=CommentRepo
type CommentRepo interface {
	SaveComment(comment *Comment) error
	GetAllCommentByAppId(appId string) ([]*Comment, error)
}
