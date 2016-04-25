package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gragas/go-sdl2/sdl"
	"github.com/gragas/jabberwock-lib/attributes"
	"github.com/gragas/jabberwock-lib/inventory"
	"github.com/gragas/jabberwock-lib/protocol"
	"github.com/gragas/jabberwock-server/serverutils"
	"net"
	"strconv"
)

type Direction byte

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Entity interface {
	GetID() uint64
	SetID(ID uint64)
	GetName() string
	SetName(name string)
	GetHealth() attributes.Health
	SetHealth(health attributes.Health)
	GetMaxHealth() attributes.Health
	SetMaxHealth(health attributes.Health)
	GetEnergy() attributes.Energy
	SetEnergy(energy attributes.Energy)
	GetMaxEnergy() attributes.Energy
	SetMaxEnergy(energy attributes.Energy)
	GetSpirit() attributes.Spirit
	SetSpirit(spirit attributes.Spirit)
	GetMaxSpirit() attributes.Spirit
	SetMaxSpirit(spirit attributes.Spirit)
	GetSummoning() attributes.Summoning
	SetSummoning(summoning attributes.Summoning)
	GetAlteration() attributes.Alteration
	SetAlteration(alteration attributes.Alteration)
	GetWillpower() attributes.Willpower
	SetWillpower(willpower attributes.Willpower)
	GetDivinity() attributes.Divinity
	SetDivinity(divinity attributes.Divinity)
	GetLifebringing() attributes.Lifebringing
	SetLifebringing(lifebringing attributes.Lifebringing)
	GetX() float32
	SetX(x float32)
	GetY() float32
	SetY(y float32)
	GetXV() float32
	SetXV(xv float32)
	GetYV() float32
	SetYV(yv float32)
	GetW() float32
	SetW(w float32)
	GetH() float32
	SetH(h float32)
	GetMovingUp() bool
	SetMovingUp(movingUp bool)
	GetMovingDown() bool
	SetMovingDown(movingDown bool)
	GetMovingRight() bool
	SetMovingRight(movingRight bool)
	GetMovingLeft() bool
	SetMovingLeft(movingLeft bool)
	GetBaseSpeed() float32
	SetBaseSpeed(baseSpeed float32)
	GetInventory() *inventory.Inventory
	GetEquipped() *([20]*inventory.Item)
	Update()
	FromBytes(bytes []byte) error
}

type EntityView interface {
	Draw(dest *sdl.Surface)
	GetObject() Entity
	SetObject(e Entity)
	GetSurface() *sdl.Surface
	SetSurface(surf *sdl.Surface)
	GetRect() *sdl.Rect
	SetRect(r *sdl.Rect)
}

func Equip(e Entity, item inventory.Item, slotNumber int) bool {
	if item.SlotNumber != slotNumber {
		return false
	}
	if item.SlotNumber == RingSlot || item.SlotNumber == PiercingSlot {
		offset := 0
		for e.GetEquipped()[item.SlotNumber+offset] != nil && offset < 4 {
			offset++
		}
		if offset == 4 {
			return false
		}
		*e.GetEquipped()[item.SlotNumber+offset] = item
		return true
	}
	if e.GetEquipped()[item.SlotNumber] != nil {
		return false
	}
	*e.GetEquipped()[item.SlotNumber] = item
	return true
}

func Unequip(e Entity, slotNumber int) (*inventory.Item, bool) {
	if e.GetEquipped()[slotNumber] != nil {
		itemPtr := e.GetEquipped()[slotNumber]
		e.GetEquipped()[slotNumber] = nil
		return itemPtr, true
	}
	return nil, false
}

func AddToInventory(e Entity, itemStack inventory.ItemStack, position int) bool {
	return e.GetInventory().AddItemStack(itemStack, position)
}

func RemoveFromInventory(e Entity, position int) (*inventory.ItemStack, bool) {
	return e.GetInventory().RemoveItemStack(position)
}

/*
 * These values are used to access the equipped member
 * of entities. E.g.,
 * entity.Equipped[HeadSlot] = $SOME_ITEM
 */
const (
	HeadSlot = iota
	LeftShoulderSlot
	RightShoulderSlot
	LeftArmSlot
	RightArmSlot
	LeftHandSlot
	RightHandSlot
	TorsoSlot
	LegsSlot
	LeftFootSlot
	RightFootSlot
	RingSlot
	PiercingSlot
	NotEquippableSlot
)

func Bytes(e Entity) []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return bytes
}

func String(e Entity) string {
	return string(Bytes(e))
}

func ShortString(e Entity) string {
	return fmt.Sprintf("<%T %v>", e, e.GetID())
}

func IDString(e Entity) string {
	IDString := strconv.FormatUint(e.GetID(), 10)
	numUnderscores := 20 - len(IDString) // 20 is the max number of digits of a uint64
	for i := 0; i < numUnderscores; i++ {
		IDString = "_" + IDString
	}
	return IDString
}

func FromIDString(str string) (uint64, error) {
	for i, c := range str {
		if c != '_' {
			num, err := strconv.ParseUint(str[i:], 10, 64)
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}
	return 0, errors.New("Could not parse ID from string!\n")
}

func StartMoveLocal(e Entity, dir Direction) {
	switch dir {
	case Up:
		e.SetYV(-e.GetBaseSpeed())
		e.SetMovingUp(true)
	case Down:
		e.SetYV(e.GetBaseSpeed())
		e.SetMovingDown(true)
	case Right:
		e.SetXV(e.GetBaseSpeed())
		e.SetMovingRight(true)
	case Left:
		e.SetXV(-e.GetBaseSpeed())
		e.SetMovingLeft(true)
	default:
		panic(errors.New("ERROR: Invalid direction.\n"))
	}
}

func StopMoveLocal(e Entity, dir Direction) {
	switch dir {
	case Up:
		e.SetMovingUp(false)
		if !e.GetMovingUp() && !e.GetMovingDown() {
			e.SetYV(0)
		}
	case Down:
		e.SetMovingDown(false)
		if !e.GetMovingUp() && !e.GetMovingDown() {
			e.SetYV(0)
		}
	case Right:
		e.SetMovingRight(false)
		if !e.GetMovingLeft() && !e.GetMovingRight() {
			e.SetXV(0)
		}
	case Left:
		e.SetMovingLeft(false)
		if !e.GetMovingLeft() && !e.GetMovingRight() {
			e.SetXV(0)
		}
	default:
		panic(errors.New("ERROR: Invalid direction.\n"))
	}
}

func MoveNet(e Entity, conn net.Conn, start bool, dir Direction) {
	var msg string
	if start {
		msg = string(protocol.EntityStartMove)
	} else {
		msg = string(protocol.EntityStopMove)
	}
	msg += string(byte(dir)) + IDString(e) + string(protocol.EndOfMessage)
	fmt.Fprintf(conn, string(msg))
}

func NewDefaultEntityView(e Entity) (*sdl.Surface, *sdl.Rect) {
	surf, err := sdl.CreateRGBSurface(0, int32(e.GetW()), int32(e.GetH()),
		32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	rect := sdl.Rect{int32(e.GetX()), int32(e.GetY()),
		int32(e.GetW()), int32(e.GetH())}
	surf.FillRect(nil, uint32(0x55555555))
	return surf, &rect
}

func Update(e Entity) {
	e.SetX(e.GetX() + e.GetXV() * serverutils.Delta)
	e.SetY(e.GetY() + e.GetYV() * serverutils.Delta)
}
