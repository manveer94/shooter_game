package gol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	gridSize = 10
)

func CreateGrid(renderer *sdl.Renderer) error {

	log.Println("creating grid")
	renderer.SetDrawColor(0, 0, 0, 255)

	//vertical
	for i := 1; i < 1000/gridSize; i++ {
		coordinate := int32(i * gridSize)
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
