package common

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	texture   *sdl.Texture
	container *Actor

	width, height         float64
	drawWidth, drawHeight float64
}

func NewSpriteRenderer(container *Actor, renderer *sdl.Renderer, filename string) *SpriteRenderer {
	tex, err := LoadTextureFromBMP(renderer, filename)
	if err != nil {
		panic(fmt.Errorf("creating sprite renderer: %v", err))
	}
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	return &SpriteRenderer{
		texture:    tex,
		container:  container,
		width:      float64(width),
		height:     float64(height),
		drawWidth:  container.DrawWidth,
		drawHeight: container.DrawHeight,
	}
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	return DrawTexture(sr.texture, sr.container.Position, sr.container.Angle, sr.drawWidth, sr.drawHeight, renderer)
}

func (sr *SpriteRenderer) OnUpdate() error {
	return nil
}

func (sr *SpriteRenderer) OnCollision(other *Actor) error {
	return nil
}
