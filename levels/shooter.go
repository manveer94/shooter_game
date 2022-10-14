package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	shooter_components2 "manveer/exp/components/shooter"
)

func CreateShooterLevel(renderer *sdl.Renderer, window *sdl.Window) *Level {

	level := Level{
		Name:     "Shooter",
		renderer: renderer,
		window:   window,
		onStart: func(renderer *sdl.Renderer) {
			shooter_components2.Actors = append(shooter_components2.Actors, shooter_components2.NewPlayer(renderer))
			createEnemies(renderer)
			shooter_components2.InitBulletPool(renderer)
		},
		onFrameChange: func(renderer *sdl.Renderer) {
			var err error
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			for _, actor := range shooter_components2.Actors {
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
			if err := shooter_components2.CheckCollisions(); err != nil {
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

	y := 10.0 + shooter_components2.Configs.BasicEnemySize/2
	for i := 0; i < 4; i++ {
		x := 45.0 + shooter_components2.Configs.BasicEnemySize/2
		for j := 0; j < 5; j++ {
			enemy := shooter_components2.NewBasicEnemy(renderer, shooter_components2.Vector{
				X: x,
				Y: y,
			})
			x = x + shooter_components2.Configs.BasicEnemySize + xMargin
			enemy.Active = true
			shooter_components2.Actors = append(shooter_components2.Actors, enemy)
		}
		y = y + shooter_components2.Configs.BasicEnemySize + yMargin
	}

}
