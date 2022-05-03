package domain

type CommentStore interface {
	Get(postId string) ([]*Comment, error)
	Create(comment *Comment) (*Comment, error)
	Delete(id string) error
	DeleteAll() error
}
