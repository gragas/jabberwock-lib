package protocol

import (
	"errors"
)

type Code byte

const (
	Register Code = iota
	Success
	BadMessageError
	GenerateEntityID
	EntityStartMoveRight
	EntityStartMoveLeft
	EntityStartMoveUp
	EntityStartMoveDown
	EntityStopMoveRight
	EntityStopMoveLeft
	EntityStopMoveUp
	EntityStopMoveDown
	UpdatePlayers
)

const EndOfMessage Code = '$'

func (c Code) String() string {
	switch c {
	case Register:
		return "Register"
	case Success:
		return "Success"
	case BadMessageError:
		return "BadMessageError"
	case GenerateEntityID:
		return "GenerateEntityID"
	case EntityStartMoveRight:
		return "EntityStartMoveRight"
	case EntityStartMoveLeft:
		return "EntityStartMoveLeft"
	case EntityStartMoveUp:
		return "EntityStartMoveUp"
	case EntityStartMoveDown:
		return "EntityStartMoveDown"
	case EntityStopMoveRight:
		return "EntityStopMoveRight"
	case EntityStopMoveLeft:
		return "EntityStopMoveLeft"
	case EntityStopMoveUp:
		return "EntityStopMoveUp"
	case EntityStopMoveDown:
		return "EntityStopMoveDown"
	case UpdatePlayers:
		return "UpdatePlayers"
	default:
		panic(errors.New("Invalid protocol.Code.\n"))
	}
}
