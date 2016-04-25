package imgutils

import (
	"github.com/gragas/go-sdl2/sdl"
	"github.com/gragas/go-sdl2/sdl_image"
)

func TextureFromImage(r *sdl.Renderer, path string) (*sdl.Texture, error) {
	s, err := img.Load(path)
	if err != nil {
		 return nil, err
	}
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return nil, err
	}
	return t, nil
}
