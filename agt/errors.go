package agt

const (
	ErrorAlreadyVoted  = 1
	ErrorVoterNotFound = 2
	ErrorDeadline      = 3
)

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}
