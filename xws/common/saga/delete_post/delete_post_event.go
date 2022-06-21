package delete_post

type DeletePostCommandType int8

const (
	DeletePost DeletePostCommandType = iota
	UnknownCommand
)

type DeletePostCommand struct {
	PostId string
	Type   DeletePostCommandType
}

type DeletePostReplyType int8

const (
	PostDeleted DeletePostReplyType = iota
	UnknownReply
)

type DeletePostReply struct {
	PostId string
	Type   DeletePostReplyType
}
