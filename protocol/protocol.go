package protocol

import (
	"errors"
)

type Code byte

const (
	Register Code = iota
	Success
	GenerateEntityID
)

const EndOfMessage Code = '$'

func (c Code) String() string {
	switch c {
	case Register:
		return "Register"
	case Success:
		return "Success"
	default:
		panic(errors.New("Invalid protocol.Code.\n"))
	}
}
