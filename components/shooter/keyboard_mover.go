package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

type keyboardMover struct {
	container *common.Actor
	speed     float64
	sr        *common.SpriteRenderer
}

func newKeyboardMover(container *common.Actor, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.GetComponent(&common.SpriteRenderer{}).(*common.SpriteRenderer),
	}
}

func (mover *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	container := mover.container
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if container.Position.X > 0 {
			container.Position.X -= mover.speed * common.Delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if container.Position.X < Configs.ScreenWidth {
			container.Position.X += mover.speed * common.Delta
		}
	}
	if keys[sdl.SCANCODE_A] == 1 {
		if container.Angle > -45 {
			container.Angle -= Configs.PlayerRotationSpeed * common.Delta
		}
	} else if keys[sdl.SCANCODE_D] == 1 {
		if container.Angle < 45 {
			container.Angle += Configs.PlayerRotationSpeed * common.Delta
		}
	}

	return nil
}

func (mover *keyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) OnCollision(other *common.Actor) error {
	return nil
}
