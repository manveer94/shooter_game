package shooter

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	texture   *sdl.Texture
	container *Actor

	width, height         float64
	drawWidth, drawHeight float64
}

func newSpriteRenderer(container *Actor, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex, err := loadTextureFromBMP(renderer, filename)
	if err != nil {
		panic(fmt.Errorf("creating sprite renderer: %v", err))
	}
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	return &spriteRenderer{
		texture:    tex,
		container:  container,
		width:      float64(width),
		height:     float64(height),
		drawWidth:  container.drawWidth,
		drawHeight: container.drawHeight,
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	return drawTexture(sr.texture, sr.container.position, sr.container.angle, sr.drawWidth, sr.drawHeight, renderer)
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func (sr *spriteRenderer) onCollision(other *Actor) error {
	return nil
}
