package api

import (
	"context"
	auth "dislinkt/common/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/promote_job"
	"dislinkt/profile_service/application"
)

type PromoteJobCommandHandler struct {
	profileService    *application.ProfileService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewPromoteJobCommandHandler(profileService *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) (*PromoteJobCommandHandler, error) {
	o := &PromoteJobCommandHandler{
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

func (handler *PromoteJobCommandHandler) handle(command *events.PromoteJobCommand) {
	reply := &events.PromoteJobReply{
		Username: command.Username,
		JobOffer: command.JobOffer,
		Profile:  command.Profile,
		Type:     events.UnknownReply,
	}
	switch command.Type {
	case events.GetProfileByToken:
		profile, err := handler.profileService.GetByToken(context.TODO(), command.Token)
		if err != nil {
			return
		}
		reply.Profile = auth.Profile{
			Username: profile.Username,
		}
		reply.Type = events.FoundProfileByToken
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
