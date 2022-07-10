package api

import (
	"context"
	events "dislinkt/common/saga/create_profile"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/security_service/application"
	securityDomain "dislinkt/security_service/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateProfileCommandHandler struct {
	securityService   *application.SecurityService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateProfileCommandHandler(securityService *application.SecurityService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateProfileCommandHandler, error) {
	o := &CreateProfileCommandHandler{
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

func (handler *CreateProfileCommandHandler) handle(command *events.CreateProfileCommand) {
	reply := &events.CreateProfileReply{Type: events.UnknownReply}
	switch command.Type {
	case events.SendVerificationEmail:
		logger := log.WithFields(logrus.Fields{
			"userId": command.Profile.Id.Hex(),
		})
		token, err := handler.securityService.GenerateVerificationToken(context.TODO())
		if err != nil {
			handler.securityService.Delete(context.TODO(), command.Profile.Id)
			logger.Errorf("CVTF: %v", err)
			reply.Type = events.ProfileNotCreated
			break
		}
		userVerification, err := handler.securityService.CreateUserVerification(context.TODO(), &securityDomain.UserVerification{
			Id:          primitive.NewObjectID(),
			Username:    command.Profile.Username,
			Token:       token,
			TimeCreated: time.Now(),
			IsVerified:  false,
		})
		if err != nil {
			handler.securityService.Delete(context.TODO(), command.Profile.Id)
			logger.Errorf("CUVF: %v", err)
			reply.Type = events.ProfileNotCreated
			break
		}
		err = handler.securityService.SendVerificationEmail(context.TODO(), command.Profile.Username, command.Profile.Email, userVerification.Token)
		if err != nil {
			logger.Errorf("SVEF: %v", err)
			handler.securityService.Delete(context.TODO(), command.Profile.Id)
			reply.Type = events.ProfileNotCreated
		}
		break
	case events.RollbackCreatedProfile:
		err := handler.securityService.Delete(context.TODO(), command.Profile.Id)
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
