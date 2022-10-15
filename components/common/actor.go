package common

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Vector struct {
	X, Y float64
}

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Actor) error
}

type Actor struct {
	Position   Vector
	Angle      float64
	Active     bool
	components []Component
	Collisions []Circle
	Tag        string
	DrawWidth  float64
	DrawHeight float64
}

func (a *Actor) Draw(renderer *sdl.Renderer) error {
	for _, comp := range a.components {
		err := comp.OnDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) Update() error {
	for _, comp := range a.components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) Collision(other *Actor) error {
	for _, comp := range a.components {
		err := comp.OnCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) AddComponent(new Component) {
	for _, existing := range a.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add a new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	a.components = append(a.components, new)
}

func (a *Actor) GetComponent(withType Component) Component {
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
