package player

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gragas/go-sdl2/sdl"
	"github.com/gragas/jabberwock-lib/attributes"
	"github.com/gragas/jabberwock-lib/entity"
	"github.com/gragas/jabberwock-lib/inventory"
	"github.com/gragas/jabberwock-lib/protocol"
)

const (
	DefaultID           = uint64(protocol.GenerateEntityID)
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
	DefaultX            = 20.0
	DefaultY            = 20.0
	DefaultXV           = 0.0
	DefaultYV           = 0.0
	DefaultW            = 25
	DefaultH            = 75
	DefaultBaseSpeed    = 0.85
)

type Player struct {
	ID                                             uint64
	Name                                           string
	Health, MaxHealth                              attributes.Health
	Energy, MaxEnergy                              attributes.Energy
	Spirit, MaxSpirit                              attributes.Spirit
	Summoning                                      attributes.Summoning
	Alteration                                     attributes.Alteration
	Willpower                                      attributes.Willpower
	Divinity                                       attributes.Divinity
	Lifebringing                                   attributes.Lifebringing
	X, Y, XV, YV, W, H                             float32
	BaseSpeed                                      float32
	MovingUp, MovingDown, MovingLeft, MovingRight  bool
	Inventory                                      *inventory.Inventory
	Equipped                                       *([20]*inventory.Item)
}

type PlayerView struct {
	PlayerPtr *Player
	Surface   *sdl.Surface
	Rect      *sdl.Rect
}

func (p *Player) Update() {
	entity.Update(p)
	// then do more stuff maybe
}

func (p *PlayerView) Draw(dest *sdl.Surface) {
	p.GetRect().X = int32(p.GetObject().GetX())
	p.GetRect().Y = int32(p.GetObject().GetY())
	p.GetSurface().Blit(nil, dest, p.GetRect())
}

func (p *PlayerView) GetObject() entity.Entity { return p.PlayerPtr }
func (p *PlayerView) SetObject(e entity.Entity) {
	switch e.(type) {
	case *Player:
		p.PlayerPtr = e.(*Player)
	default:
		panic(errors.New(fmt.Sprintf("Cannot set a PlayerView's PlayerPtr equal to a %T", e)))
	}
}

func (p *Player) String() string {
	return entity.String(p)
}

func (p *PlayerView) GetSurface() *sdl.Surface { return p.Surface }
func (p *PlayerView) SetSurface(surf *sdl.Surface) { p.Surface = surf }
func (p *PlayerView) GetRect() *sdl.Rect { return p.Rect }
func (p *PlayerView) SetRect(r *sdl.Rect) { p.Rect = r }

func (p *Player) GetID() uint64 { return p.ID }
func (p *Player) SetID(ID uint64) { p.ID = ID }
func (p *Player) GetName() string { return p.Name }
func (p *Player) SetName(Name string) { p.Name = Name }
func (p *Player) GetHealth() attributes.Health { return p.Health }
func (p *Player) SetHealth(health attributes.Health) { p.Health = health }
func (p *Player) GetEnergy() attributes.Energy { return p.Energy }
func (p *Player) SetEnergy(energy attributes.Energy) { p.Energy = energy }
func (p *Player) GetSpirit() attributes.Spirit { return p.Spirit }
func (p *Player) SetSpirit(spirit attributes.Spirit) { p.Spirit = spirit }
func (p *Player) GetMaxHealth() attributes.Health { return p.MaxHealth }
func (p *Player) SetMaxHealth(maxHealth attributes.Health) { p.MaxHealth = maxHealth }
func (p *Player) GetMaxEnergy() attributes.Energy { return p.MaxEnergy }
func (p *Player) SetMaxEnergy(maxEnergy attributes.Energy) { p.MaxEnergy = maxEnergy }
func (p *Player) GetMaxSpirit() attributes.Spirit { return p.MaxSpirit }
func (p *Player) SetMaxSpirit(maxSpirit attributes.Spirit) { p.MaxSpirit = maxSpirit }
func (p *Player) GetSummoning() attributes.Summoning { return p.Summoning }
func (p *Player) SetSummoning(summoning attributes.Summoning) { p.Summoning = summoning }
func (p *Player) GetAlteration() attributes.Alteration { return p.Alteration }
func (p *Player) SetAlteration(alteration attributes.Alteration) { p.Alteration = alteration }
func (p *Player) GetWillpower() attributes.Willpower { return p.Willpower }
func (p *Player) SetWillpower(willpower attributes.Willpower) { p.Willpower = willpower }
func (p *Player) GetDivinity() attributes.Divinity { return p.Divinity }
func (p *Player) SetDivinity(divinity attributes.Divinity) { p.Divinity = divinity }
func (p *Player) GetLifebringing() attributes.Lifebringing { return p.Lifebringing }
func (p *Player) SetLifebringing(lifebringing attributes.Lifebringing) { p.Lifebringing = lifebringing }
func (p *Player) GetX() float32 { return p.X }
func (p *Player) SetX(x float32) { p.X = x }
func (p *Player) GetY() float32 { return p.Y }
func (p *Player) SetY(y float32) { p.Y = y }
func (p *Player) GetXV() float32 { return p.XV }
func (p *Player) SetXV(xv float32) { p.XV = xv }
func (p *Player) GetYV() float32 { return p.YV }
func (p *Player) SetYV(yv float32) { p.YV = yv }
func (p *Player) GetW() float32 { return p.W }
func (p *Player) SetW(w float32) { p.W = w }
func (p *Player) GetH() float32 { return p.H }
func (p *Player) SetH(h float32) { p.H = h }
func (p *Player) GetMovingUp() bool { return p.MovingUp }
func (p *Player) SetMovingUp(movingUp bool) { p.MovingUp = movingUp }
func (p *Player) GetMovingDown() bool { return p.MovingDown }
func (p *Player) SetMovingDown(movingDown bool) { p.MovingDown = movingDown }
func (p *Player) GetMovingLeft() bool { return p.MovingLeft }
func (p *Player) SetMovingLeft(movingLeft bool) { p.MovingLeft = movingLeft }
func (p *Player) GetMovingRight() bool { return p.MovingRight }
func (p *Player) SetMovingRight(movingRight bool) { p.MovingRight = movingRight }
func (p *Player) GetBaseSpeed() float32 { return p.BaseSpeed }
func (p *Player) SetBaseSpeed(baseSpeed float32) { p.BaseSpeed = baseSpeed }
func (p *Player) GetInventory() *inventory.Inventory { return p.Inventory }
func (p *Player) GetEquipped() *([20]*inventory.Item) { return p.Equipped }

func NewDefaultPlayer() *Player {
	p := Player{ID: DefaultID,
		Name: DefaultName,
		Health: DefaultHealth,
		MaxHealth: DefaultMaxHealth,
		Energy: DefaultEnergy,
		MaxEnergy: DefaultMaxEnergy,
		Spirit: DefaultSpirit,
		MaxSpirit: DefaultMaxSpirit,
		Summoning: DefaultSummoning,
		Alteration: DefaultAlteration,
		Willpower: DefaultWillpower,
		Divinity: DefaultDivinity,
		Lifebringing: DefaultLifebringing,
		X: DefaultX, Y: DefaultY, XV: DefaultXV, YV: DefaultYV,
		W: DefaultW, H: DefaultH,
		BaseSpeed: DefaultBaseSpeed}
	return &p
}

func (p *Player) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, &p)
}
