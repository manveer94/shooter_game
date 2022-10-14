package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/shooter"
	"manveer/exp/levels"
)

var window *sdl.Window
var renderer *sdl.Renderer

func main() {
	initWindow()
	defer func(window *sdl.Window) {
		log.Println("closing window")
		err := window.Destroy()
		if err != nil {
			log.Fatalln("error while closing window", err)
		}
	}(window)
	initRenderer()
	defer func(renderer *sdl.Renderer) {
		log.Println("destroying renderer")
		err := renderer.Destroy()
		if err != nil {
			log.Fatalln("error while destroying renderer", err)
		}
	}(renderer)

	shooterLevel := levels.CreateShooterLevel(renderer, window)
	shooterLevel.Start()

	//gameOfLife := levels.CreateGameOfLife(renderer, window)
	//gameOfLife.Start()

}

func initWindow() {
	var err error
	log.Println("initializing SDL")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalln("error while initializing SDL", err)
	}

	log.Println("Creating window")
	window, err = sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(shooter.Configs.ScreenWidth), int32(shooter.Configs.ScreenHeight),
		sdl.WINDOW_OPENGL)

	if err != nil {
		log.Fatalln("error while initializing window", err)
	}

}

func initRenderer() {
	var err error
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	log.Println("initializing renderer")
	if err != nil {
		log.Fatalln("error while initializing the renderer")
	}

}
