package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components"
	"time"
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

	initEventLoop()
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
		int32(components.Configs.ScreenWidth), int32(components.Configs.ScreenHeight),
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
		log.Fatalln("error while initializing the rederer")
	}

}

func initEventLoop() {
	var err error
	components.Actors = append(components.Actors, components.NewPlayer(renderer))
	createEnemies()
	components.InitBulletPool(renderer)

	for {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				log.Println("quit event received")
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, actor := range components.Actors {
			if actor.Active {
				err = actor.Update()
				if err != nil {
					log.Fatalf("updating actor: %v", err)
				}
				err = actor.Draw(renderer)
				if err != nil {
					log.Fatalf("drawing actor: %v", err)
				}
			}

		}
		if err := components.CheckCollisions(); err != nil {
			log.Fatalf("checking collisions:%v \n", err)
		}
		renderer.Present()
		components.Delta = time.Since(frameStartTime).Seconds() * components.Configs.TargetTicksPerSecond
	}
}

func createEnemies() {
	log.Println("Creating enemies")
	xMargin := 70.0
	yMargin := 50.0

	y := 10.0 + components.Configs.BasicEnemySize/2
	for i := 0; i < 4; i++ {
		x := 45.0 + components.Configs.BasicEnemySize/2
		for j := 0; j < 5; j++ {
			enemy := components.NewBasicEnemy(renderer, components.Vector{
				X: x,
				Y: y,
			})
			x = x + components.Configs.BasicEnemySize + xMargin
			enemy.Active = true
			components.Actors = append(components.Actors, enemy)
		}
		y = y + components.Configs.BasicEnemySize + yMargin
	}

}
