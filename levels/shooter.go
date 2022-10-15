package levels

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/shooter"
)

func CreateShooterLevel(renderer *sdl.Renderer, window *sdl.Window) *Level {

	level := Level{
		Name:     "Shooter",
		renderer: renderer,
		window:   window,
		onStart: func(renderer *sdl.Renderer) error {
			shooter.Actors = append(shooter.Actors, shooter.NewPlayer(renderer))
			createEnemies(renderer)
			shooter.InitBulletPool(renderer)
			return nil
		},
		onFrameChange: func(renderer *sdl.Renderer) error {
			var err error
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			for _, actor := range shooter.Actors {
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
			if err := shooter.CheckCollisions(); err != nil {
				log.Fatalf("checking collisions:%v \n", err)
			}
			renderer.Present()
			return nil
		},
	}
	return &level
}

func createEnemies(renderer *sdl.Renderer) {
	log.Println("Creating enemies")
	xMargin := 70.0
	yMargin := 50.0

	y := 10.0 + shooter.Configs.BasicEnemySize/2
	for i := 0; i < 4; i++ {
		x := 45.0 + shooter.Configs.BasicEnemySize/2
		for j := 0; j < 5; j++ {
			enemy := shooter.NewBasicEnemy(renderer, shooter.Vector{
				X: x,
				Y: y,
			})
			x = x + shooter.Configs.BasicEnemySize + xMargin
			enemy.Active = true
			shooter.Actors = append(shooter.Actors, enemy)
		}
		y = y + shooter.Configs.BasicEnemySize + yMargin
	}

}
