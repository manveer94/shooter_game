package gol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

func createCell(gridCoordinate *common.VectorI, state bool) (*common.Actor, error) {
	x := gridCoordinate.X
	y := gridCoordinate.Y

	position := &common.Vector{X: float64(x) * gridSize, Y: float64(y) * gridSize}
	size := &common.Vector{X: gridSize, Y: gridSize}

	cell := &common.Actor{
		Active: state,
		Tag:    fmt.Sprintf("cell%d-%d", x, y),
	}
	//cell.AddComponent(&cellLifeCycle{
	//	container:      cell,
	//	gridCoordinate: gridCoordinate,
	//	alive:          state,
	//})
	cell.AddComponent(&common.RectRenderer{
		Fill:         true,
		Position:     position,
		Size:         size,
		FillColor:    &common.Color{R: 170, G: 170, B: 170, A: 255},
		Outline:      true,
		OutlineColor: &common.Color{R: 70, G: 70, B: 70, A: 255},
	})

	return cell, nil
}

type cellLifeCycle struct {
	container      *common.Actor
	gridCoordinate *common.VectorI
	alive          bool
}

func (cell *cellLifeCycle) OnUpdate() error {
	return nil
}

func (cell *cellLifeCycle) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (cell *cellLifeCycle) OnCollision(other *common.Actor) error {
	return nil
}
