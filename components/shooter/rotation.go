package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
)

type rotation struct {
	container *Actor
	speed     float64
}

func newSimpleRotation(container *Actor, speed float64) *rotation {
	return &rotation{
		container: container,
		speed:     speed,
	}
}

func (r *rotation) onUpdate() error {
	r.container.angle += r.speed * Delta
	if r.container.angle == 360 {
		r.container.angle = 0
	}
	return nil
}

func (r *rotation) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (r *rotation) onCollision(other *Actor) error {
	return nil
}
