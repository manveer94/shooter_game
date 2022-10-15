package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/shooter"
	"time"
)

type Level struct {
	Name          string
	window        *sdl.Window
	renderer      *sdl.Renderer
	onStart       func(*sdl.Renderer) error
	onFrameChange func(*sdl.Renderer) error
	onStop        func() error
}

func (lv Level) Start() {
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
		err := lv.onFrameChange(lv.renderer)
		if err != nil {
			log.Printf("fatal error occured exiting game: %v", err)
			return
		}
		shooter.Delta = time.Since(frameStartTime).Seconds() * shooter.Configs.TargetTicksPerSecond
	}

}
