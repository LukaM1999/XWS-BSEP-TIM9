package update_profile

import "dislinkt/common/domain"

type UpdateProfileCommandType int8

const (
	UpdateProfile UpdateProfileCommandType = iota
	RollbackUpdatedProfile
	UnknownCommand
)

type UpdateProfileCommand struct {
	Profile      domain.Profile
	OldUsername  string
	OldFirstName string
	OldLastName  string
	OldIsPrivate bool
	Type         UpdateProfileCommandType
}

type UpdateProfileReplyType int8

const (
	ProfileUpdated UpdateProfileReplyType = iota
	ProfileNotUpdated
	ProfileUpdateRolledBack
	UnknownReply
)

type UpdateProfileReply struct {
	Profile      domain.Profile
	OldUsername  string
	OldFirstName string
	OldLastName  string
	OldIsPrivate bool
	Type         UpdateProfileReplyType
}
