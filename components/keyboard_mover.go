package components

import "github.com/veandco/go-sdl2/sdl"

type keyboardMover struct {
	container *Actor
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *Actor, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	container := mover.container
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if container.position.X > 0 {
			container.position.X -= mover.speed * Delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if container.position.X < Configs.ScreenWidth {
			container.position.X += mover.speed * Delta
		}
	}
	if keys[sdl.SCANCODE_A] == 1 {
		if container.angle > -45 {
			container.angle -= Configs.PlayerRotationSpeed * Delta
		}
	} else if keys[sdl.SCANCODE_D] == 1 {
		if container.angle < 45 {
			container.angle += Configs.PlayerRotationSpeed * Delta
		}
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *Actor) error {
	return nil
}
