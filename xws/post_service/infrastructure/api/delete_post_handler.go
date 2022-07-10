package api

import (
	"context"
	events "dislinkt/common/saga/delete_post"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/post_service/application"
)

type DeletePostCommandHandler struct {
	postService       *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeletePostCommandHandler(postService *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeletePostCommandHandler, error) {
	o := &DeletePostCommandHandler{
		postService:       postService,
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
		var err = handler.postService.Delete(context.TODO(), command.PostId)
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
