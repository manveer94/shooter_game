package gol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

const (
	gridSize   = 10
	screenSize = 1000
	gridLength = screenSize / gridSize
)

func Initialize(renderer *sdl.Renderer) error {
	createCellMatrix()
	background, err := CreateBackground(renderer)
	if err != nil {
		return fmt.Errorf("creating background: %v", err)
	}
	common.Actors = append(common.Actors, background)
	return nil
}
