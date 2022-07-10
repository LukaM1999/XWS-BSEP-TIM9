package api

import (
	"context"
	"dislinkt/comment_service/application"
	events "dislinkt/common/saga/delete_post"
	saga "dislinkt/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeletePostCommandHandler struct {
	commentService    *application.CommentService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeletePostCommandHandler(commentService *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeletePostCommandHandler, error) {
	o := &DeletePostCommandHandler{
		commentService:    commentService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeletePostCommandHandler) handle(command *events.DeletePostCommand) {
	reply := &events.DeletePostReply{
		PostId: command.PostId,
		Type:   events.UnknownReply,
	}
	switch command.Type {
	case events.DeletePost:
		id, err := primitive.ObjectIDFromHex(command.PostId)
		if err != nil {
			return
		}
		err = handler.commentService.DeletePostComments(context.TODO(), id)
		if err != nil {
			return
		}
		reply.Type = events.PostDeleted
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
