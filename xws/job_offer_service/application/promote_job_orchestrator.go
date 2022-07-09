package application

import (
	auth "dislinkt/common/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/promote_job"
	"dislinkt/job_offer_service/domain"
)

type PromoteJobOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewPromoteJobOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*PromoteJobOrchestrator, error) {
	o := &PromoteJobOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *PromoteJobOrchestrator) Start(token string, username string, jobOffer domain.JobOffer) error {
	event := &events.PromoteJobCommand{
		Username: username,
		JobOffer: auth.JobOffer(jobOffer),
		Token:    token,
		Type:     events.GetProfileByToken,
	}
	return o.commandPublisher.Publish(event)
}

func (o *PromoteJobOrchestrator) handle(reply *events.PromoteJobReply) {
	command := events.PromoteJobCommand{
		JobOffer: reply.JobOffer,
		Token:    "",
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *PromoteJobOrchestrator) nextCommandType(reply events.PromoteJobReplyType) events.PromoteJobCommandType {
	switch reply {
	case events.FoundProfileByToken:
		return events.PromoteJob
	default:
		return events.UnknownCommand
	}
}
