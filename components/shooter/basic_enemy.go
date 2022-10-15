package shooter

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
)

type BasicEnemy struct {
	common.Actor
}

func NewBasicEnemy(renderer *sdl.Renderer, position common.Vector) *common.Actor {
	e := &common.Actor{}
	e.Angle = 180
	e.Position = position
	e.DrawWidth = Configs.BasicEnemySize
	e.DrawHeight = Configs.BasicEnemySize
	//sr := NewSpriteRenderer(e, renderer, "sprites/enemy.bmp")
	//e.AddComponent(sr)

	idleSequence, err := common.NewSequence("sprites/basic_enemy/idle", 3, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle Sequence: %v", err))
	}

	destroySequence, err := common.NewSequence("sprites/basic_enemy/destroy", 25, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy Sequence: %v", err))
	}

	sequences := map[string]*common.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := common.NewAnimator(e, sequences, "idle")

	e.AddComponent(animator)

	e.AddComponent(common.NewSimpleRotation(e, 0.5))

	e.AddComponent(newVulnerableToBullets(e))

	e.Collisions = append(e.Collisions, common.Circle{
		Radius: Configs.BasicEnemySize / 2,
		Center: position,
	})
	return e
}
