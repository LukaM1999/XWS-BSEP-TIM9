package api

import (
	"context"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_profile"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
)

type UpdateProfileCommandHandler struct {
	postService       *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(postService *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
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
		newProfile := &domain.Profile{
			Id:        command.Profile.Id,
			FirstName: command.Profile.FirstName,
			LastName:  command.Profile.LastName,
		}
		err := handler.postService.UpdateProfile(context.TODO(), newProfile.Id, newProfile)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		oldProfile := &domain.Profile{
			Id:        command.Profile.Id,
			FirstName: command.OldFirstName,
			LastName:  command.OldLastName,
		}
		err := handler.postService.UpdateProfile(context.TODO(), oldProfile.Id, oldProfile)
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
