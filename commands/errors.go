package commands

import "errors"

var (
	ErrRecordAlreadyStarted = errors.New("record already started")
	ErrRecordNotFound       = errors.New("record not found")
	ErrNoActiveRecord       = errors.New("no active record")

	ErrUnknownCommand = errors.New("unknown command")
)
