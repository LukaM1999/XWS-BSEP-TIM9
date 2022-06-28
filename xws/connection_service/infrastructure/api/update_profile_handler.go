package api

import (
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_profile"
	"dislinkt/connection_service/application"
)

type UpdateProfileCommandHandler struct {
	connectionService *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(connectionService *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		connectionService: connectionService,
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
		if command.Profile.IsPrivate == command.OldIsPrivate {
			return
		}
		err := handler.connectionService.UpdatePrivacy(command.Profile.Id.Hex())
		if err != nil {
			reply.Type = events.ProfileNotUpdated
			break
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		if command.Profile.IsPrivate == command.OldIsPrivate {
			return
		}
		err := handler.connectionService.UpdatePrivacy(command.Profile.Id.Hex())
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
