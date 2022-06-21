package api

import (
	events "dislinkt/common/saga/create_profile"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/connection_service/application"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProfileCommandHandler struct {
	connectionService *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateProfileCommandHandler(connectionService *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateProfileCommandHandler, error) {
	o := &CreateProfileCommandHandler{
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

func (handler *CreateProfileCommandHandler) handle(command *events.CreateProfileCommand) {
	reply := &events.CreateProfileReply{
		Profile: command.Profile,
		Type:    events.UnknownReply,
	}
	switch command.Type {
	case events.CreateProfile:
		profilePrivacy := &domain.ProfilePrivacy{
			Id:        primitive.NewObjectID(),
			UserId:    command.Profile.Id,
			IsPrivate: false,
		}
		_, err := handler.connectionService.CreateProfilePrivacy(profilePrivacy)
		if err != nil {
			reply.Type = events.ProfileNotCreated
			break
		}
		//reply.Type = events.ProfileCreated
		break
	case events.RollbackCreatedProfile:
		err := handler.connectionService.DeleteProfilePrivacy(command.Profile.Id)
		if err != nil {
			return
		}
		reply.Type = events.ProfileCreationRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
