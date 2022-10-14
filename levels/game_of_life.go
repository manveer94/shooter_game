package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func CreateGameOfLife(renderer *sdl.Renderer, window *sdl.Window) *Level {
	level := &Level{}
	level.Name = "Game of life"
	level.renderer = renderer
	level.window = window

	level.onStart = func(renderer *sdl.Renderer) {
		window.SetSize(1000, 1000)
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawLine(10, 10, 100, 100)
		//renderer.Clear()

		renderer.Present()
	}

	level.onFrameChange = func(renderer *sdl.Renderer) {

	}

	return level
}
