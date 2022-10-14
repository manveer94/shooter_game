package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/shooter_components"
)

func CreateShooterLevel(renderer *sdl.Renderer, window *sdl.Window) *Level {

	level := Level{
		Name:     "Shooter",
		renderer: renderer,
		window:   window,
		onStart: func(renderer *sdl.Renderer) {
			shooter_components.Actors = append(shooter_components.Actors, shooter_components.NewPlayer(renderer))
			createEnemies(renderer)
			shooter_components.InitBulletPool(renderer)
		},
		onFrameChange: func(renderer *sdl.Renderer) {
			var err error
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			for _, actor := range shooter_components.Actors {
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
			if err := shooter_components.CheckCollisions(); err != nil {
				log.Fatalf("checking collisions:%v \n", err)
			}
			renderer.Present()
		},
	}
	return &level
}

func createEnemies(renderer *sdl.Renderer) {
	log.Println("Creating enemies")
	xMargin := 70.0
	yMargin := 50.0

	y := 10.0 + shooter_components.Configs.BasicEnemySize/2
	for i := 0; i < 4; i++ {
		x := 45.0 + shooter_components.Configs.BasicEnemySize/2
		for j := 0; j < 5; j++ {
			enemy := shooter_components.NewBasicEnemy(renderer, shooter_components.Vector{
				X: x,
				Y: y,
			})
			x = x + shooter_components.Configs.BasicEnemySize + xMargin
			enemy.Active = true
			shooter_components.Actors = append(shooter_components.Actors, enemy)
		}
		y = y + shooter_components.Configs.BasicEnemySize + yMargin
	}

}
