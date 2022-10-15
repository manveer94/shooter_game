package common

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Rotation struct {
	container *Actor
	speed     float64
}

func NewSimpleRotation(container *Actor, speed float64) *Rotation {
	return &Rotation{
		container: container,
		speed:     speed,
	}
}

func (r *Rotation) OnUpdate() error {
	r.container.Angle += r.speed * Delta
	if r.container.Angle == 360 {
		r.container.Angle = 0
	}
	return nil
}

func (r *Rotation) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (r *Rotation) OnCollision(other *Actor) error {
	return nil
}
