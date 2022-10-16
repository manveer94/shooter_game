package gol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

func CreateBackground(renderer *sdl.Renderer) (*common.Actor, error) {
	bg := &common.Actor{
		Active: true,
		Tag:    "background",
	}
	bg.AddComponent(&grid{gridSize: gridSize})
	bg.AddComponent(&onLeftClick{})
	bg.AddComponent(&simulate{})
	return bg, nil
}

type grid struct {
	gridSize int
}

func (g *grid) OnDraw(renderer *sdl.Renderer) error {
	renderer.SetDrawColor(200, 200, 200, 50)
	for i := 1; i < 1000/g.gridSize; i++ {
		coordinate := int32(i * g.gridSize)
		err := renderer.DrawLine(coordinate, 0, coordinate, 1000)
		if err != nil {
			return fmt.Errorf("drawing vertical grid lines: %v", err)
		}
		err = renderer.DrawLine(0, coordinate, 1000, coordinate)
		if err != nil {
			return fmt.Errorf("drawing horizontal grid lines: %v", err)
		}
	}
	return nil
}

func (g *grid) OnUpdate() error {
	return nil
}

func (g *grid) OnCollision(other *common.Actor) error {
	return nil
}
