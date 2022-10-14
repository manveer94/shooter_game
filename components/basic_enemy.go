package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type BasicEnemy struct {
	Actor
}

func NewBasicEnemy(renderer *sdl.Renderer, position Vector) *Actor {
	e := &Actor{}
	e.angle = 180
	e.position = position
	e.drawWidth = Configs.BasicEnemySize
	e.drawHeight = Configs.BasicEnemySize
	//sr := newSpriteRenderer(e, renderer, "sprites/enemy.bmp")
	//e.addComponent(sr)

	idleSequence, err := newSequence("sprites/basic_enemy/idle", 3, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}

	destroySequence, err := newSequence("sprites/basic_enemy/destroy", 25, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}

	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := newAnimator(e, sequences, "idle")

	e.addComponent(animator)

	e.addComponent(newSimpleRotation(e, 0.5))

	e.addComponent(newVulnerableToBullets(e))

	e.collisions = append(e.collisions, circle{
		radius: Configs.BasicEnemySize / 2,
		center: position,
	})
	return e
}
