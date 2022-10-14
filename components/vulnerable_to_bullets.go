package components

import "github.com/veandco/go-sdl2/sdl"

type vulnerableToBullets struct {
	container *Actor
	animator  *animator
}

func newVulnerableToBullets(container *Actor) *vulnerableToBullets {
	return &vulnerableToBullets{
		container: container,
		animator:  container.getComponent(&animator{}).(*animator),
	}
}

func (v *vulnerableToBullets) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (v *vulnerableToBullets) onUpdate() error {
	if v.animator.finished && v.animator.current == "destroy" {
		v.container.Active = false
	}
	return nil
}

func (v *vulnerableToBullets) onCollision(other *Actor) error {
	if other.tag == "bullet" {
		//v.container.Active = false
		v.animator.setSequence("destroy")
		//v.animator.
	}
	return nil
}
