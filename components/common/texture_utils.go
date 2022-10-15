package common

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func DrawTexture(texture *sdl.Texture, position Vector, rotation float64, drawWidth float64, drawHeight float64, renderer *sdl.Renderer) error {
	_, _, width, height, err := texture.Query()
	//Converting actor coordinates to top left of sprite

	x := position.X - drawWidth/2.0
	y := position.Y - drawHeight/2.0

	err = renderer.CopyEx(
		texture,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(drawWidth), H: int32(drawHeight)},
		rotation,
		&sdl.Point{X: int32(drawWidth / 2), Y: int32(drawHeight / 2)},
		sdl.FLIP_NONE)

	return err
}

func LoadTextureFromBMP(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	surface, err := sdl.LoadBMP(filename)
	if err != nil {
		log.Printf("not able to load actor sprite from path %s: %v\n", filename, err)
		return nil, err
	}
	defer surface.Free()
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Printf("not able create actor texture: %v\n", err)
		return nil, err

	}
	return texture, nil
}
