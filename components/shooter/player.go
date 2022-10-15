package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"manveer/exp/components/common"
	"time"
)

func NewPlayer(renderer *sdl.Renderer) *common.Actor {
	player := &common.Actor{}
	player.Position = common.Vector{
		X: Configs.ScreenWidth / 2,
		Y: Configs.ScreenHeight - Configs.PlayerSize/2,
	}

	player.DrawWidth = Configs.PlayerSize
	player.DrawHeight = Configs.PlayerSize

	player.Active = true
	sr := common.NewSpriteRenderer(player, renderer, "sprites/ship.bmp")
	player.AddComponent(sr)

	mover := newKeyboardMover(player, Configs.PlayerBaseSpeed)
	player.AddComponent(mover)

	shooter := newKeyboardShooter(player, Configs.PlayerShotCoolDown)
	player.AddComponent(shooter)
	return player
}

type keyboardShooter struct {
	container *common.Actor
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *common.Actor, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (shooter *keyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	position := shooter.container.Position
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
		bul.Position.X = x
		bul.Position.Y = y
		bul.Angle = shooter.container.Angle
	} else {
		log.Println("no bullets left to shoot")
	}
}

func (shooter *keyboardShooter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) OnCollision(other *common.Actor) error {
	return nil
}
