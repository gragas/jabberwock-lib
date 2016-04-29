package textures

import (
	"github.com/gragas/go-sdl2/sdl"	
)

var Textures map[string]*sdl.Texture

func Init() {
	Textures = make(map[string]*sdl.Texture)
}
