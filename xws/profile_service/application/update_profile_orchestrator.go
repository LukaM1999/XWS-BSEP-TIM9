package application

import (
	"dislinkt/common/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_profile"
)

type UpdateProfileOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUpdateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileOrchestrator, error) {
	o := &UpdateProfileOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UpdateProfileOrchestrator) Start(profile *domain.Profile, oldUsername string,
	oldFirstName string, oldLastName string, oldIsPrivate bool) error {
	event := &events.UpdateProfileCommand{
		Profile:      *profile,
		OldUsername:  oldUsername,
		OldFirstName: oldFirstName,
		OldLastName:  oldLastName,
		OldIsPrivate: oldIsPrivate,
		Type:         events.UpdateProfile,
	}
	return o.commandPublisher.Publish(event)
}

func (o *UpdateProfileOrchestrator) handle(reply *events.UpdateProfileReply) {
	command := events.UpdateProfileCommand{
		Profile:      reply.Profile,
		OldUsername:  reply.OldUsername,
		OldFirstName: reply.OldFirstName,
		OldLastName:  reply.OldLastName,
		OldIsPrivate: reply.OldIsPrivate,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *UpdateProfileOrchestrator) nextCommandType(reply events.UpdateProfileReplyType) events.UpdateProfileCommandType {
	switch reply {
	case events.ProfileNotUpdated:
		return events.RollbackUpdatedProfile
	default:
		return events.UnknownCommand
	}
}
