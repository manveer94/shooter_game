package levels

import (
	"manveer/exp/components/gol"
)

func CreateGameOfLife() *Level {
	level := &Level{}
	level.Name = "Game of life"
	level.onStart = gol.Initialize
	level.windowWidth = 1000
	level.windowHeight = 1000
	return level
}
