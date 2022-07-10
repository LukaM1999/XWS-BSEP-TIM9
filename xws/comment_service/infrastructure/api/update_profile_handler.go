package api

import (
	"context"
	"dislinkt/comment_service/application"
	"dislinkt/comment_service/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_profile"
)

type UpdateProfileCommandHandler struct {
	commentService    *application.CommentService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(commentService *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
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

func (handler *UpdateProfileCommandHandler) handle(command *events.UpdateProfileCommand) {
	reply := &events.UpdateProfileReply{
		Profile:      command.Profile,
		Type:         events.UnknownReply,
		OldUsername:  command.OldUsername,
		OldFirstName: command.OldFirstName,
		OldLastName:  command.OldLastName,
		OldIsPrivate: command.OldIsPrivate,
	}
	switch command.Type {
	case events.UpdateProfile:
		if command.Profile.FirstName == command.OldFirstName && command.Profile.LastName == command.OldLastName {
			return
		}
		newProfile := &domain.CommentCreator{
			Id:        command.Profile.Id,
			FirstName: command.Profile.FirstName,
			LastName:  command.Profile.LastName,
		}
		err := handler.commentService.UpdateCommentCreator(context.TODO(), newProfile.Id, newProfile)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		oldProfile := &domain.CommentCreator{
			Id:        command.Profile.Id,
			FirstName: command.OldFirstName,
			LastName:  command.OldLastName,
		}
		err := handler.commentService.UpdateCommentCreator(context.TODO(), oldProfile.Id, oldProfile)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdateRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
