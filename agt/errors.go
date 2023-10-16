package agt

import "net/http"

const (
	ErrorAlreadyVoted  = http.StatusForbidden
	ErrorVoterNotFound = http.StatusBadRequest
	ErrorDeadline      = http.StatusServiceUnavailable
)

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) StatusCode() int {
	return e.Code
}
