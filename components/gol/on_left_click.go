package gol

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/common"
)

var isPressed2 = false

type onLeftClick struct {
	container *common.Actor
	isPressed bool
}

func (o *onLeftClick) OnUpdate() error {
	x, y, state := sdl.GetMouseState()

	if !o.isPressed && state&sdl.ButtonLMask() != 0 {
		o.isPressed = true
	}

	if o.isPressed && state&sdl.ButtonLMask() == 0 {
		o.isPressed = false
		log.Printf("left click: %v,%v\n", x/gridSize, y/gridSize)

		coordinate := &common.VectorI{
			X: x / gridSize, Y: y / gridSize,
		}
		if isCellAlive(coordinate) {
			killCell(coordinate)
		} else {
			resurrectCell(coordinate)
		}
	}

	//if !isPressed2 && state&sdl.ButtonRMask() != 0 {
	//	isPressed2 = true
	//}
	//if isPressed2 && state&sdl.ButtonRMask() == 0 {
	//	isPressed2 = false
	//}

	return nil
}

func (o *onLeftClick) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (o *onLeftClick) OnCollision(other *common.Actor) error {
	return nil
}
