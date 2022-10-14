package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

func NewPlayer(renderer *sdl.Renderer) *Actor {
	player := &Actor{}
	player.position = Vector{
		X: Configs.ScreenWidth / 2,
		Y: Configs.ScreenHeight - Configs.PlayerSize/2,
	}

	player.drawWidth = Configs.PlayerSize
	player.drawHeight = Configs.PlayerSize

	player.Active = true
	sr := newSpriteRenderer(player, renderer, "sprites/ship.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, Configs.PlayerBaseSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, Configs.PlayerShotCoolDown)
	player.addComponent(shooter)
	return player
}

type keyboardShooter struct {
	container *Actor
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *Actor, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()
	position := shooter.container.position
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			shooter.shoot(position.X-20.0, position.Y)
			shooter.shoot(position.X+20.0, position.Y)
			shooter.lastShot = time.Now()
		}
	}
	return nil
}

func (shooter *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.Active = true
		bul.position.X = x
		bul.position.Y = y
		bul.angle = shooter.container.angle
	} else {
		log.Println("no bullets left to shoot")
	}
}

func (shooter *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) onCollision(other *Actor) error {
	return nil
}
