package player

import (
	"github.com/gragas/go-sdl2/sdl"
	"github.com/gragas/jabberwock-lib/entity"
)

const (
	DefaultName         = "Default Name"
	DefaultHealth       = 10.0
	DefaultMaxHealth    = 10.0
	DefaultEnergy       = 10.0
	DefaultMaxEnergy    = 10.0
	DefaultSpirit       = 10.0
	DefaultMaxSpirit    = 10.0
	DefaultSummoning    = 5.0
	DefaultAlteration   = 5.0
	DefaultWillpower    = 5.0
	DefaultDivinity     = 5.0
	DefaultLifebringing = 5.0
	DefaultX            = 0.0
	DefaultY            = 0.0
	DefaultXv           = 0.0
	DefaultYv           = 0.0
	DefaultW            = 25
	DefaultH            = 75
)

type Player struct {
	entity.Entity
}

type PlayerView struct {
	Player
	Surface *sdl.Surface
}
