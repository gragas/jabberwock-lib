package entity

import (
	"encoding/json"
	"github.com/gragas/jabberwock-lib/attributes"
	"github.com/gragas/jabberwock-lib/inventory"
)

type Entity struct {
	ID                 uint64
	Name               string
	Health, MaxHealth  attributes.Health
	Energy, MaxEnergy  attributes.Energy
	Spirit, MaxSpirit  attributes.Spirit
	Summoning          attributes.Summoning
	Alteration         attributes.Alteration
	Willpower          attributes.Willpower
	Divinity           attributes.Divinity
	Lifebringing       attributes.Lifebringing
	X, Y, Xv, Yv, W, H float32
	Inventory          inventory.Inventory
	Equipped           [20]*inventory.Item
}

func (entity Entity) Equip(item inventory.Item, slotNumber int) bool {
	if item.SlotNumber != slotNumber {
		return false
	}
	if item.SlotNumber == RingSlot || item.SlotNumber == PiercingSlot {
		offset := 0
		for entity.Equipped[item.SlotNumber+offset] != nil && offset < 4 {
			offset++
		}
		if offset == 4 {
			return false
		}
		*entity.Equipped[item.SlotNumber+offset] = item
		return true
	}
	if entity.Equipped[item.SlotNumber] != nil {
		return false
	}
	*entity.Equipped[item.SlotNumber] = item
	return true
}

func (entity Entity) Unequip(slotNumber int) (*inventory.Item, bool) {
	if entity.Equipped[slotNumber] != nil {
		itemPtr := entity.Equipped[slotNumber]
		entity.Equipped[slotNumber] = nil
		return itemPtr, true
	}
	return nil, false
}

func (entity Entity) AddToInventory(itemStack inventory.ItemStack, position int) bool {
	return entity.Inventory.AddItemStack(itemStack, position)
}

func (entity Entity) RemoveFromInventory(position int) (*inventory.ItemStack, bool) {
	return entity.Inventory.RemoveItemStack(position)
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

func (e Entity) Bytes() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (e Entity) String() string {
	return string(e.Bytes())
}

func FromBytes(bytes []byte) (*Entity, error) {
	var e Entity
	err := json.Unmarshal(bytes, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
