package components

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Bullet struct {
	Actor
}

var bulletCount = 0

func NewBullet(renderer *sdl.Renderer) *Actor {
	b := &Actor{}
	b.drawWidth = Configs.BulletWidth
	b.drawHeight = Configs.BulletHeight
	sr := newSpriteRenderer(b, renderer, "sprites/player_bullet.bmp")
	b.addComponent(sr)
	b.addComponent(newBulletTrajectory(b, Configs.BulletSpeed))
	b.Active = false
	b.tag = "bullet"
	col := circle{
		center: b.position,
		radius: Configs.BulletWidth,
	}
	b.collisions = append(b.collisions, col)
	return b
}

var BulletPool []*Actor

func InitBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullet := NewBullet(renderer)
		BulletPool = append(BulletPool, bullet)
		Actors = append(Actors, bullet)
	}
}

func bulletFromPool() (*Actor, bool) {
	for _, bul := range BulletPool {
		if !bul.Active {
			return bul, true
		}
	}
	return nil, false
}

type bulletTrajectory struct {
	container *Actor
	speed     float64
}

func newBulletTrajectory(container *Actor, speed float64) *bulletTrajectory {
	return &bulletTrajectory{
		container: container,
		speed:     speed,
	}
}

func (b *bulletTrajectory) onUpdate() error {
	container := b.container
	container.position.Y -= Configs.BulletSpeed * math.Cos(container.angle*(math.Pi/180)) * Delta
	container.position.X += Configs.BulletSpeed * math.Sin(container.angle*(math.Pi/180)) * Delta
	if container.position.Y < 0 {
		container.Active = false
	}
	b.container.collisions[0].center = container.position

	return nil
}

func (b *bulletTrajectory) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (b *bulletTrajectory) onCollision(other *Actor) error {
	b.container.Active = false
	return nil
}
