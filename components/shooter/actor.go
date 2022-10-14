package shooter

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Vector struct {
	X, Y float64
}

type Component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollision(other *Actor) error
}

type Actor struct {
	position   Vector
	angle      float64
	Active     bool
	components []Component
	collisions []circle
	tag        string
	drawWidth  float64
	drawHeight float64
}

func (a *Actor) Draw(renderer *sdl.Renderer) error {
	for _, comp := range a.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) Update() error {
	for _, comp := range a.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Actor) Collision(other *Actor) error {
	for _, comp := range a.components {
		err := comp.onCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) addComponent(new Component) {
	for _, existing := range a.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add a new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	a.components = append(a.components, new)
}

func (a *Actor) getComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range a.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf(
		"no component with type %v",
		reflect.TypeOf(withType)))
}

var Actors []*Actor
