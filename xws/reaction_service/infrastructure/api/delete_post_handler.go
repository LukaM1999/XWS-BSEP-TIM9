package api

import (
	events "dislinkt/common/saga/delete_post"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/reaction_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeletePostCommandHandler struct {
	reactionService   *application.ReactionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeletePostCommandHandler(reactionService *application.ReactionService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeletePostCommandHandler, error) {
	o := &DeletePostCommandHandler{
		reactionService:   reactionService,
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
		err = handler.reactionService.DeletePostReactions(id)
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
