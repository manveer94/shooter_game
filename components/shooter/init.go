package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/common"
)

func createEnemies(renderer *sdl.Renderer) {
	log.Println("Creating enemies")
	xMargin := 70.0
	yMargin := 50.0

	y := 10.0 + Configs.BasicEnemySize/2
	for i := 0; i < 4; i++ {
		x := 45.0 + Configs.BasicEnemySize/2
		for j := 0; j < 5; j++ {
			enemy := NewBasicEnemy(renderer, common.Vector{
				X: x,
				Y: y,
			})
			x = x + Configs.BasicEnemySize + xMargin
			enemy.Active = true
			common.Actors = append(common.Actors, enemy)
		}
		y = y + Configs.BasicEnemySize + yMargin
	}
}

func Initialize(renderer *sdl.Renderer) error {
	common.Actors = append(common.Actors, NewPlayer(renderer))
	createEnemies(renderer)
	InitBulletPool(renderer)
	return nil
}
