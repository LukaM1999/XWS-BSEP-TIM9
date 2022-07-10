package api

import (
	"context"
	auth "dislinkt/common/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/promote_job"
	"dislinkt/job_offer_service/application"
	"dislinkt/job_offer_service/domain"
)

type PromoteJobCommandHandler struct {
	jobOfferService   *application.JobOfferService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewPromoteJobCommandHandler(jobOfferService *application.JobOfferService, publisher saga.Publisher, subscriber saga.Subscriber) (*PromoteJobCommandHandler, error) {
	o := &PromoteJobCommandHandler{
		jobOfferService:   jobOfferService,
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
		Type: events.UnknownReply,
	}
	switch command.Type {
	case events.PromoteJob:
		if command.Username != command.Profile.Username {
			reply.Type = events.UnknownReply
			break
		}
		job, err := handler.jobOfferService.CreateJob(context.TODO(), (*domain.JobOffer)(&command.JobOffer))
		if err != nil {
			return
		}
		reply.JobOffer = auth.JobOffer(*job)
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
