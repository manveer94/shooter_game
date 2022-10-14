package components

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

//func newSpriteRendererWithSize(container *Actor, renderer *sdl.Renderer, filename string, drawWidth float64, drawHeight float64) *spriteRenderer {
//	tex := loadTextureFromBMP(renderer, filename)
//	_, _, width, height, err := tex.Query()
//	if err != nil {
//		panic(fmt.Errorf("querying texture: %v", err))
//	}
//	return &spriteRenderer{
//		texture:    tex,
//		container:  container,
//		width:      float64(width),
//		height:     float64(height),
//		drawWidth:  drawWidth,
//		drawHeight: drawHeight,
//	}
//}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	//Converting actor coordinates to top left of sprite

	return drawTexture(sr.texture, sr.container.position, sr.container.angle, sr.drawWidth, sr.drawHeight, renderer)
	//width := int32(sr.width)
	//height := int32(sr.height)
	//x := sr.container.position.X - sr.drawWidth/2.0
	//y := sr.container.position.Y - sr.drawHeight/2.0
	//
	//err := renderer.CopyEx(
	//	sr.texture,
	//	&sdl.Rect{X: 0, Y: 0, W: width, H: height},
	//	&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.drawWidth), H: int32(sr.drawHeight)},
	//	sr.container.angle,
	//	&sdl.Point{X: int32(sr.drawWidth / 2), Y: int32(sr.drawHeight / 2)},
	//	sdl.FLIP_NONE)
	//
	//return err
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func (sr *spriteRenderer) onCollision(other *Actor) error {
	return nil
}
