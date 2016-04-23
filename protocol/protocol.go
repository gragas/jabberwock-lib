package protocol

import (
	"errors"
)

type Code byte

const (
	Register Code = iota
	Success
	Handshake
	GenerateEntityID
	EntityStartMove
	EntityStopMove
	UpdatePlayers
)

const EndOfMessage Code = '$'

func (c Code) String() string {
	switch c {
	case Register:
		return "Register"
	case Success:
		return "Success"
	case Handshake:
		return "Handshake"
	case GenerateEntityID:
		return "GenerateEntityID"
	case EntityStartMove:
		return "EntityStartMove"
	case EntityStopMove:
		return "EntityStopMove"
	case UpdatePlayers:
		return "UpdatePlayers"
	default:
		panic(errors.New("Invalid protocol.Code.\n"))
	}
}
