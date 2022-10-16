package gol

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/common"
	"time"
)

var cellsAddr [][]*common.Actor
var cellsState [][]bool
var tempCellsState [][]bool
var simulationRunning = false
var lastUpdate = time.Now()

func UpdateMatrix() error {

	if simulationRunning {
		if time.Since(lastUpdate) < time.Millisecond*100 {
			return nil
		}
		lastUpdate = time.Now()
		for x := 0; x < gridLength; x++ {
			for y := 0; y < gridLength; y++ {
				tempCellsState[x][y] = cellsState[x][y]
			}
		}
		for x := 0; x < gridLength; x++ {
			for y := 0; y < gridLength; y++ {
				coordinate := &common.VectorI{
					X: int32(x), Y: int32(y),
				}
				totalNeighbours := countNeighbours(x, y)
				if isCellAlive(coordinate) {
					if totalNeighbours < 2 {
						killCell(coordinate)
					}
					if totalNeighbours > 3 {
						killCell(coordinate)
					}
				} else {
					if totalNeighbours == 3 {
						resurrectCell(coordinate)
					}
				}
				//if x == 0 && y == 0 {
				//	fmt.Println(totalNeighbours)
				//}
			}
		}
		//simulationRunning = false
	}
	return nil
}

func countNeighbours(x, y int) int {
	totalNeighbours := checkNeighbourAlive(x, y, -1, -1) + checkNeighbourAlive(x, y, 0, -1) + checkNeighbourAlive(x, y, 1, -1) + checkNeighbourAlive(x, y, -1, 0)
	totalNeighbours = totalNeighbours + checkNeighbourAlive(x, y, 1, 0) + checkNeighbourAlive(x, y, -1, 1) + checkNeighbourAlive(x, y, 0, 1) + checkNeighbourAlive(x, y, 1, 1)
	return totalNeighbours
}

func checkNeighbourAlive(x, y, dx, dy int) int {
	alive := 0
	maxGridIndex := gridLength - 1
	x = x + dx
	y = y + dy

	if x < 0 {
		x = maxGridIndex
	}
	if x > maxGridIndex {
		x = 0
	}

	if y < 0 {
		y = maxGridIndex
	}
	if y > maxGridIndex {
		y = 0
	}
	if tempCellsState[x][y] {
		alive = 1
	}
	return alive
}

type simulate struct {
	spacePressed bool
}

func (s *simulate) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	if !s.spacePressed && keys[sdl.SCANCODE_SPACE] == 1 {
		s.spacePressed = true
	}

	if s.spacePressed && keys[sdl.SCANCODE_SPACE] == 0 {
		s.spacePressed = false
		simulationRunning = !simulationRunning
		if simulationRunning {
			log.Println("simulation started")
		}

		if !simulationRunning {
			log.Println("simulation stopped")
		}
	}

	return nil
}

func (s *simulate) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (s *simulate) OnCollision(other *common.Actor) error {
	return nil
}

func isCellAlive(coordinate *common.VectorI) bool {
	return cellsState[coordinate.X][coordinate.Y]
}

func killCell(coordinate *common.VectorI) {
	cellsState[coordinate.X][coordinate.Y] = false
	cellsAddr[coordinate.X][coordinate.Y].Active = false

}

func resurrectCell(coordinate *common.VectorI) {
	cellsState[coordinate.X][coordinate.Y] = true
	cellsAddr[coordinate.X][coordinate.Y].Active = true
}

func createCellMatrix() {

	cellsAddr = make([][]*common.Actor, gridLength)
	cellsState = make([][]bool, gridLength)
	tempCellsState = make([][]bool, gridLength)
	for i, _ := range cellsAddr {
		cellsAddr[i] = make([]*common.Actor, gridLength)
		cellsState[i] = make([]bool, gridLength)
		tempCellsState[i] = make([]bool, gridLength)
		for j := 0; j < gridLength; j++ {
			cellsAddr[i][j], _ = createCell(&common.VectorI{X: int32(i), Y: int32(j)}, false)
			common.Actors = append(common.Actors, cellsAddr[i][j])
		}
	}
}
