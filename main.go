package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/levels"
)

var window *sdl.Window
var renderer *sdl.Renderer

func main() {

	//shooterLevel := levels.CreateShooterLevel()
	//shooterLevel.Start()

	gameOfLife := levels.CreateGameOfLife()
	gameOfLife.Start()

}
