package common

import (
	"github.com/veandco/go-sdl2/sdl"
)

type RectRenderer struct {
	Size         *Vector
	Position     *Vector
	FillColor    *Color
	OutlineColor *Color
	Fill         bool
	Outline      bool
}

func (rect *RectRenderer) OnUpdate() error {
	return nil
}

func (rect *RectRenderer) OnDraw(renderer *sdl.Renderer) error {
	if rect.Fill {
		if &rect.FillColor == nil {
			rect.FillColor = &Color{R: 0, G: 0, B: 0, A: 255}
		}
		renderer.SetDrawColor(rect.FillColor.R, rect.FillColor.G, rect.FillColor.B, rect.FillColor.A)
		renderer.FillRect(&sdl.Rect{
			X: int32(rect.Position.X), Y: int32(rect.Position.Y),
			W: int32(rect.Size.X), H: int32(rect.Size.Y),
		})
	}

	if rect.Outline {
		if &rect.Outline == nil {
			rect.OutlineColor = &Color{R: 0, G: 0, B: 0, A: 255}
		}
		renderer.SetDrawColor(rect.OutlineColor.R, rect.OutlineColor.G, rect.OutlineColor.B, rect.OutlineColor.A)
		renderer.DrawRect(&sdl.Rect{
			X: int32(rect.Position.X), Y: int32(rect.Position.Y),
			W: int32(rect.Size.X), H: int32(rect.Size.Y),
		})
	}
	return nil
}

func (rect *RectRenderer) OnCollision(other *Actor) error {
	return nil
}
