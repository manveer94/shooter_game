package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/gol"
)

func CreateGameOfLife(renderer *sdl.Renderer, window *sdl.Window) *Level {
	level := &Level{}
	level.Name = "Game of life"
	level.renderer = renderer
	level.window = window

	level.onStart = func(renderer *sdl.Renderer) error {
		window.SetSize(1000, 1000)
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		err := gol.CreateGrid(renderer)
		if err != nil {
			return err
		}

		renderer.Present()
		return nil
	}

	//isPressed := false

	level.onFrameChange = func(renderer *sdl.Renderer) error {

		//x, y, state := sdl.GetMouseState()
		////fmt.Printf("%d, %d, %v\n", x, y, state)
		//
		//if state&sdl.ButtonLMask() != 0 {
		//	renderer.SetDrawColor(255, 0, 0, 255)
		//	isPressed = true
		//}
		//
		//if state&sdl.ButtonMMask() != 0 {
		//	renderer.SetDrawColor(0, 255, 0, 255)
		//	isPressed = true
		//}
		//
		//if state&sdl.ButtonRMask() != 0 {
		//	renderer.SetDrawColor(0, 0, 255, 255)
		//	isPressed = true
		//}
		//
		//if state == sdl.RELEASED && isPressed {
		//	isPressed = false
		//
		//	fmt.Printf("%d, %d, %v\n", x, y, state&sdl.ButtonLMask())
		//	renderer.FillRect(&sdl.Rect{
		//		X: x, Y: y,
		//		W: 50, H: 50,
		//	})
		//
		//	renderer.Present()
		//}
		return nil
	}

	return level
}
