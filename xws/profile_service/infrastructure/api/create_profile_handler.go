package api

import (
	"context"
	events "dislinkt/common/saga/create_profile"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/profile_service/application"
	"dislinkt/profile_service/domain"
	"time"
)

type CreateProfileCommandHandler struct {
	profileService    *application.ProfileService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateProfileCommandHandler(profileService *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateProfileCommandHandler, error) {
	o := &CreateProfileCommandHandler{
		profileService:    profileService,
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
	}
	switch command.Type {
	case events.CreateProfile:
		profile := &domain.Profile{
			Id:             command.Profile.Id,
			Username:       command.Profile.Username,
			FirstName:      command.Profile.FirstName,
			LastName:       command.Profile.LastName,
			FullName:       command.Profile.FullName,
			DateOfBirth:    time.Time{},
			PhoneNumber:    "",
			Email:          command.Profile.Email,
			Gender:         "",
			IsPrivate:      false,
			Biography:      "",
			Education:      nil,
			WorkExperience: nil,
			Skills:         nil,
			Interests:      nil,
			AgentToken:     "",
		}
		err := handler.profileService.Create(context.TODO(), profile)
		if err != nil {
			reply.Type = events.ProfileNotCreated
			break
		}
		reply.Type = events.ProfileCreated
		break
	case events.RollbackCreatedProfile:
		err := handler.profileService.Delete(context.TODO(), command.Profile.Id.Hex())
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
