package application

import (
	"dislinkt/comment_service/domain"
)

type CommentService struct {
	store domain.CommentStore
}

func NewCommentService(store domain.CommentStore) *CommentService {
	return &CommentService{
		store: store,
	}
}

func (service *CommentService) Get(postId string) ([]*domain.Comment, error) {
	return service.store.Get(postId)
}

func (service *CommentService) Create(comment *domain.Comment) (*domain.Comment, error) {
	return service.store.Create(comment)
}

func (service *CommentService) Delete(id string) error {
	return service.store.Delete(id)
}
