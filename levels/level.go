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
	onStart       func(*sdl.Renderer)
	onFrameChange func(*sdl.Renderer)
	onStop        func() error
}

func (lv Level) Start() {
	log.Printf("initializing %s level\n", lv.Name)
	lv.onStart(lv.renderer)
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
		lv.onFrameChange(lv.renderer)
		shooter.Delta = time.Since(frameStartTime).Seconds() * shooter.Configs.TargetTicksPerSecond
	}

}
