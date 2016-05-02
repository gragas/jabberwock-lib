package inventory

import (
	"fmt"
	"github.com/gragas/go-sdl2/sdl"
	"github.com/gragas/jabberwock-lib/imgutils"
	"github.com/gragas/jabberwock-lib/textures"
	"path/filepath"
	"strings"
)

type Quantity uint32

/********** Item **********/
type Item struct {
	Name       string
	SlotNumber int
	X, Y       float32
}

func (item Item) String() string {
	return fmt.Sprintf("%s", item.Name)
}
/***************************/

/********** ItemStack **********/
type ItemStack struct {
	Item     Item
	Quantity Quantity
}

func (itemStack ItemStack) String() string {
	return fmt.Sprintf("%v x %v", itemStack.Item, itemStack.Quantity)
}
/*******************************/

/********** ItemStack **********/
type ItemStackView struct {
	ItemStackPtr *ItemStack
	Rect *sdl.Rect
	Texture *sdl.Texture
}

const (
	DefaultItemStackViewW    = 32
	DefaultItemStackViewH    = 32
	DefaultItemStackViewPath = "jabberwock-assets" + string(filepath.Separator)
)

func (is *ItemStack) NewDefaultItemStackView(r *sdl.Renderer, x int32, y int32) *ItemStackView {
	/* Create an sdl.Rect */
	rect := &sdl.Rect{x, y, DefaultItemStackViewW, DefaultItemStackViewH}

	var texture *sdl.Texture
	filePath := DefaultItemStackViewPath + "items" + string(filepath.Separator) + is.Item.Name + ".png"
	if textures.Textures[filePath] == nil { // if the inventory texture isn't in the cache
		var err error
		texture, err = imgutils.TextureFromImage(r, filePath)
		if err != nil {
			panic(err)
		}
		textures.Textures[filePath] = texture // cache it
	} else {
		texture = textures.Textures[filePath] // otherwise, retreive it from the cache
	}
	return &ItemStackView{ItemStackPtr: is, Rect: rect, Texture: texture}
}

func (isv *ItemStackView) Render(r *sdl.Renderer) {
	// render the quantity onto the texture
}

func (isv *ItemStackView) Draw(r *sdl.Renderer) {
	r.Copy(isv.Texture, nil, isv.Rect)	
}
/*******************************/

/********** Inventory **********/
const maxInventoryItemStacks uint32 = 32

type Inventory struct {
	ItemStacks [maxInventoryItemStacks]*ItemStack
}

func (inventory Inventory) String() string {
	var itemStackNames []string
	for _, itemStack := range inventory.ItemStacks {
		itemStackNames = append(itemStackNames, itemStack.Item.Name)
	}
	return strings.Join(itemStackNames, ", ")
}

func (inventory Inventory) AddItemStack(itemStack ItemStack, position int) bool {
	if inventory.ItemStacks[position] == nil {
		*inventory.ItemStacks[position] = itemStack
		return true
	} else if inventory.ItemStacks[position].Item.Name == itemStack.Item.Name {
		(*inventory.ItemStacks[position]).Quantity += itemStack.Quantity
		return true
	}
	return false
}

func (inventory Inventory) RemoveItemStack(position int) (*ItemStack, bool) {
	if inventory.ItemStacks[position] != nil {
		itemStackPtr := inventory.ItemStacks[position]
		inventory.ItemStacks[position] = nil
		return itemStackPtr, true
	}
	return nil, false
}
/***********************************/

/********** InventoryView **********/
const (
	DefaultInventoryViewX    = 272
	DefaultInventoryViewY    = 472
	DefaultInventoryViewW    = 256
	DefaultInventoryViewH    = 128
	DefaultInventoryViewPath = "jabberwock-assets" + string(filepath.Separator) + "inventory" + string(filepath.Separator) + "inventory.png"
)

type InventoryView struct {
	InventoryPtr *Inventory
	Rect         *sdl.Rect
	Texture      *sdl.Texture
}

func (i *Inventory) NewDefaultInventoryView(r *sdl.Renderer) *InventoryView {
	/* Create an sdl.Rect */
	rect := &sdl.Rect{DefaultInventoryViewX, DefaultInventoryViewY, DefaultInventoryViewW, DefaultInventoryViewH}

	var texture *sdl.Texture
	if textures.Textures[DefaultInventoryViewPath] == nil { // if the inventory texture isn't in the cache
		var err error
		texture, err = imgutils.TextureFromImage(r, DefaultInventoryViewPath)
		if err != nil {
			panic(err)
		}
		textures.Textures[DefaultInventoryViewPath] = texture // cache it
	} else {
		texture = textures.Textures[DefaultInventoryViewPath] // otherwise, retreive it from the cache
	}
	return &InventoryView{InventoryPtr: i, Rect: rect, Texture: texture}
}

func (iv *InventoryView) Render(r *sdl.Renderer) {
	/*
	for _, itemStack := range iv.InventoryPtr.ItemStacks {
		// render each item stack
	}
        */
}

func (iv *InventoryView) Draw(r *sdl.Renderer) {
	r.Copy(iv.Texture, nil, iv.Rect)
}
/***********************************/
