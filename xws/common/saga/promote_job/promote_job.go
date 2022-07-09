package promote_job

import "dislinkt/common/domain"

type PromoteJobCommandType int8

const (
	GetProfileByToken PromoteJobCommandType = iota
	PromoteJob
	UnknownCommand
)

type PromoteJobCommand struct {
	Profile  domain.Profile
	Username string
	JobOffer domain.JobOffer
	Token    string
	Type     PromoteJobCommandType
}

type PromoteJobReplyType int8

const (
	FoundProfileByToken PromoteJobReplyType = iota
	UnknownReply
)

type PromoteJobReply struct {
	Username string
	JobOffer domain.JobOffer
	Profile  domain.Profile
	Type     PromoteJobReplyType
}
