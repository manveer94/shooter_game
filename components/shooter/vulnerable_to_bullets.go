package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

type vulnerableToBullets struct {
	container *common.Actor
	animator  *common.Animator
}

func newVulnerableToBullets(container *common.Actor) *vulnerableToBullets {
	return &vulnerableToBullets{
		container: container,
		animator:  container.GetComponent(&common.Animator{}).(*common.Animator),
	}
}

func (v *vulnerableToBullets) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (v *vulnerableToBullets) OnUpdate() error {
	if v.animator.Finished && v.animator.Current == "destroy" {
		v.container.Active = false
	}
	return nil
}

func (v *vulnerableToBullets) OnCollision(other *common.Actor) error {
	if other.Tag == "bullet" {
		//v.Container.Active = false
		v.animator.SetSequence("destroy")
		//v.Animator.
	}
	return nil
}
