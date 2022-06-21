package application

import (
	events "dislinkt/common/saga/delete_post"
	saga "dislinkt/common/saga/messaging"
)

type DeletePostOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeletePostOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeletePostOrchestrator, error) {
	o := &DeletePostOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeletePostOrchestrator) Start(postId string) error {
	event := &events.DeletePostCommand{
		PostId: postId,
		Type:   events.DeletePost,
	}
	return o.commandPublisher.Publish(event)
}

func (o *DeletePostOrchestrator) handle(reply *events.DeletePostReply) {
	command := events.DeletePostCommand{
		PostId: reply.PostId,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeletePostOrchestrator) nextCommandType(reply events.DeletePostReplyType) events.DeletePostCommandType {
	switch reply {
	case events.PostDeleted:
		return events.UnknownCommand
	default:
		return events.UnknownCommand
	}
}
