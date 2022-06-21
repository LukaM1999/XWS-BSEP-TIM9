package create_profile

import "dislinkt/common/domain"

type CreateProfileCommandType int8

const (
	CreateProfile CreateProfileCommandType = iota
	RollbackCreatedProfile
	SendVerificationEmail
	UnknownCommand
)

type CreateProfileCommand struct {
	Profile domain.Profile
	Type    CreateProfileCommandType
}

type CreateProfileReplyType int8

const (
	ProfileCreated CreateProfileReplyType = iota
	ProfileNotCreated
	ProfileCreationRolledBack
	UnknownReply
)

type CreateProfileReply struct {
	Profile domain.Profile
	Type    CreateProfileReplyType
}
