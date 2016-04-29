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
	"github.com/gragas/jabberwock-lib/imgutils"
	"github.com/gragas/jabberwock-lib/textures"
//	"math"
//	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	DefaultID            = uint64(protocol.GenerateEntityID)
	DefaultName          = "Default Name"
	DefaultHealth        = 10.0
	DefaultMaxHealth     = 10.0
	DefaultEnergy        = 10.0
	DefaultMaxEnergy     = 10.0
	DefaultSpirit        = 10.0
	DefaultMaxSpirit     = 10.0
	DefaultSummoning     = 5.0
	DefaultAlteration    = 5.0
	DefaultWillpower     = 5.0
	DefaultDivinity      = 5.0
	DefaultLifebringing  = 5.0
	DefaultX             = 20.0
	DefaultY             = 20.0
	DefaultXV            = 0.0
	DefaultYV            = 0.0
	DefaultW             = 25
	DefaultH             = 75
	DefaultBaseSpeed     = 0.02
)

const (
	DefaultPlayerViewA     = 0x55
	DefaultSpritePath      = "jabberwock-assets" + string(filepath.Separator) + "soul"
	DefaultSpriteDuration  = time.Duration(333 * time.Millisecond)
)

type Player struct {
	Id                                             uint64
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
	PlayerPtr             *Player
	Texture               *sdl.Texture
	Surface               *sdl.Surface
	Rect                  *sdl.Rect
	SpritePath            string
	SpriteTicks           time.Duration
	SpriteDuration        time.Duration
	NumStationarySprites  int
	NumMovingRightSprites int
	NumMovingLeftSprites  int
	NumMovingUpSprites    int
	NumMovingDownSprites  int
	StationarySprites     []*sdl.Texture
	MovingRightSprites    []*sdl.Texture
	MovingLeftSprites     []*sdl.Texture
	MovingUpSprites       []*sdl.Texture
	MovingDownSprites     []*sdl.Texture
}

func (p *Player) Update() {
	entity.Update(p)
	// then do more stuff maybe
}

func (p *PlayerView) Draw(r *sdl.Renderer, dest *sdl.Surface, delta time.Duration) {
	/* update the rect position */
	p.Rect.X = int32(p.PlayerPtr.X)
	p.Rect.Y = int32(p.PlayerPtr.Y)
	/****************************/

	/* udpate the SpriteTicks for animations */
	p.SpriteTicks = (p.SpriteTicks + delta) % p.SpriteDuration
	spritePercent := float32(p.SpriteTicks) / float32(p.SpriteDuration)
	player := p.PlayerPtr
	if !player.Moving() {
		p.Texture = p.StationarySprites[int(spritePercent * float32(p.NumStationarySprites))]
	} else if player.MovingRight {
		p.Texture = p.MovingRightSprites[int(spritePercent * float32(p.NumMovingRightSprites))]
	} else if player.MovingLeft {
		p.Texture = p.MovingLeftSprites[int(spritePercent * float32(p.NumMovingLeftSprites))]
	} else if player.MovingUp {
		p.Texture = p.MovingUpSprites[int(spritePercent * float32(p.NumMovingUpSprites))]
	}  else if player.MovingDown {
		p.Texture = p.MovingDownSprites[int(spritePercent * float32(p.NumMovingDownSprites))]
	}
	/*****************************************/

	/* actually blit the playerview onto the destination surface */
	r.Copy(p.Texture, nil, p.Rect)
	/*************************************************************/
}

func (p *Player) Bytes() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (p *Player) String() string {
	return string(p.Bytes())
}

func (p *Player) Moving() bool {
	return p.GetMovingRight() || p.GetMovingLeft() || p.GetMovingUp() || p.GetMovingDown()
}

func (p *Player) GetID() uint64 { return p.Id }
func (p *Player) SetID(id uint64) { p.Id = id }
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
func (p *Player) SetInventory(i *inventory.Inventory) { p.Inventory = i }
func (p *Player) GetEquipped() *([20]*inventory.Item) { return p.Equipped }
func (p *Player) SetEquipped(e *([20]*inventory.Item)) { p.Equipped = e }

func NewDefaultPlayer() *Player {
	p := Player{Id: DefaultID,
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

func isValidSpriteImage(name string) bool {
	lenName := len(name)
	if lenName < 5 {
		return false
	}
	if name[lenName-4:] != ".png" {
		return false
	}
	numString := name[:lenName-4]
	_, err := strconv.ParseUint(numString, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func (p *Player) NewDefaultPlayerView(r *sdl.Renderer) *PlayerView {
	/* Create an sdl.Rect */
	rect := sdl.Rect{int32(p.GetX()), int32(p.GetY()), int32(p.GetW()), int32(p.GetH())}

	/* Initialize slices */
	sliceCap := 5
	stationarySprites := make([]*sdl.Texture, sliceCap)
	movingRightSprites := make([]*sdl.Texture, sliceCap)
	movingLeftSprites := make([]*sdl.Texture, sliceCap)
	movingUpSprites := make([]*sdl.Texture, sliceCap)
	movingDownSprites := make([]*sdl.Texture, sliceCap)

	/* Gather the sprite images for this *PlayerView */
	baseImgPath := DefaultSpritePath
	subPaths := [...]string{"stationary", "movingRight", "movingLeft", "movingUp", "movingDown"}
	var counts [5]int
	for i, subPath := range subPaths {
		dirPath := baseImgPath + string(filepath.Separator) + subPath
		if _, err := os.Stat(dirPath); err == nil { // if the directory exists
			d, err := os.Open(dirPath); defer d.Close() //open the directory
			if err != nil { panic(err) }
			fileinfos, err := d.Readdir(-1) // get the fileinfos
			if err != nil { panic(err) }
			for _, fileinfo := range fileinfos {
				filePath := dirPath + string(filepath.Separator) + fileinfo.Name()
				if fileinfo.Mode().IsRegular() && isValidSpriteImage(fileinfo.Name()) {
					var texture *sdl.Texture
					if textures.Textures[filePath] == nil {
						texture, err = imgutils.TextureFromImage(r, filePath)
						if err != nil {
							wd, wderr := os.Getwd()
							if wderr == nil { fmt.Println("pwd:", wd) }
							fmt.Printf("CLIENT: Could not load image '%s'\n", filePath)
							panic(err)
						}
						textures.Textures[filePath] = texture // cache it
					} else {
						texture = textures.Textures[filePath] // retrieve from cache
					}
					if texture == nil {
						fmt.Printf("Failed to load texture '%s'\n", filePath)
						continue
					}
					err = texture.SetBlendMode(sdl.BLENDMODE_BLEND)
					if err != nil { panic(err) }
					if counts[i] < sliceCap {
						switch subPaths[i] {
						case "stationary":
							stationarySprites[counts[i]] = texture
						case "movingRight":
							movingRightSprites[counts[i]] = texture
						case "movingLeft":
							movingLeftSprites[counts[i]] = texture
						case "movingUp":
							movingUpSprites[counts[i]] = texture
						case "movingDown":
							movingDownSprites[counts[i]] = texture
						}
					} else {
						switch subPaths[i] {
						case "stationary":
							stationarySprites = append(stationarySprites, texture)
						case "movingRight":
							movingRightSprites = append(movingRightSprites, texture)
						case "movingLeft":
							movingLeftSprites = append(movingLeftSprites, texture)
						case "movingUp":
							movingUpSprites = append(movingUpSprites, texture)
						case "movingDown":
							movingDownSprites = append(movingDownSprites, texture)
						}
					}
					counts[i]++ // found another valid sprite image for this subPath
				}
			}		
		}
	}
	
	/* If we failed to load at least one stationary sprite */
	if counts[0] < 1 {
		panic(errors.New("Failed to load at least one stationary sprite!\n"))
	}

	/* Reset the draw blend mode of the renderer */
	err := r.SetDrawBlendMode(sdl.BLENDMODE_BLEND); if err != nil { panic(err) }

	return &PlayerView{PlayerPtr: p,
		Texture: stationarySprites[0],
		Surface: nil, Rect: &rect,
		SpritePath: DefaultSpritePath,
		SpriteTicks: time.Duration(0),
		SpriteDuration: DefaultSpriteDuration,
		NumStationarySprites: counts[0],
		NumMovingRightSprites: counts[1],
		NumMovingLeftSprites: counts[2],
		NumMovingUpSprites: counts[3],
		NumMovingDownSprites: counts[4],
		StationarySprites: stationarySprites,
		MovingRightSprites: movingRightSprites,
		MovingLeftSprites: movingLeftSprites,
		MovingUpSprites: movingUpSprites,
		MovingDownSprites: movingDownSprites}
}

func (p *Player) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, p)
}
