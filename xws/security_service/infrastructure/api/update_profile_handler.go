package api

import (
	"context"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_profile"
	"dislinkt/security_service/application"
)

type UpdateProfileCommandHandler struct {
	securityService   *application.SecurityService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(securityService *application.SecurityService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		securityService:   securityService,
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
		if command.Profile.Username == command.OldUsername {
			return
		}
		_, err := handler.securityService.Update(context.TODO(), command.Profile.Id, command.Profile.Username)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		_, err := handler.securityService.Update(context.TODO(), command.Profile.Id, command.Profile.Username)
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
