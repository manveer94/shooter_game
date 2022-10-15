package levels

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/common"
	"time"
)

type Level struct {
	Name         string
	windowWidth  float64
	windowHeight float64
	window       *sdl.Window
	renderer     *sdl.Renderer
	onStart      func(*sdl.Renderer) error
	onStop       func() error
}

func (lv *Level) Start() {

	lv.initWindow()
	lv.initRenderer()

	lv.run()
	defer func(window *sdl.Window) {
		log.Println("closing window")
		err := window.Destroy()
		if err != nil {
			log.Fatalln("error while closing window", err)
		}
	}(lv.window)

	defer func(renderer *sdl.Renderer) {
		log.Println("destroying renderer")
		err := renderer.Destroy()
		if err != nil {
			log.Fatalln("error while destroying renderer", err)
		}
	}(lv.renderer)

}

func (lv *Level) run() {
	log.Printf("initializing %s level\n", lv.Name)
	err := lv.onStart(lv.renderer)
	if err != nil {
		log.Printf("error occurred while initializing level: %v", err)
		return
	}
	log.Printf("starting %s level\n", lv.Name)
	for {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				log.Println("quit event received")
				if lv.onStop != nil {
					log.Println("Calling onStop")
					err := lv.onStop()
					if err != nil {
						log.Printf("error while executing onStop: %v\n", err)
					}
				}
				return
			}
		}
		err := lv.frameChange()
		if err != nil {
			log.Printf("fatal error occured exiting game: %v", err)
			return
		}
		common.Delta = time.Since(frameStartTime).Seconds() * common.TargetTicksPerSecond
	}
}

func (lv *Level) frameChange() error {
	var err error
	lv.renderer.SetDrawColor(255, 255, 255, 255)
	lv.renderer.Clear()

	for _, actor := range common.Actors {
		if actor.Active {
			err = actor.Update()
			if err != nil {
				return fmt.Errorf("updating actor: %v", err)
			}
			err = actor.Draw(lv.renderer)
			if err != nil {
				return fmt.Errorf("drawing actor: %v", err)
			}
		}

	}
	if err := common.CheckCollisions(); err != nil {
		return fmt.Errorf("checking collisions:%v \n", err)
	}
	lv.renderer.Present()
	return nil
}

func (lv *Level) initWindow() {
	var err error
	log.Println("initializing SDL")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalln("error while initializing SDL", err)
	}

	log.Println("Creating window")
	lv.window, err = sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(lv.windowWidth), int32(lv.windowHeight),
		sdl.WINDOW_OPENGL)

	if err != nil {
		log.Fatalln("error while initializing window", err)
	}
}

func (lv *Level) initRenderer() {
	var err error
	lv.renderer, err = sdl.CreateRenderer(lv.window, -1, sdl.RENDERER_ACCELERATED)
	log.Println("initializing renderer")
	if err != nil {
		log.Fatalln("error while initializing the renderer")
	}
}
