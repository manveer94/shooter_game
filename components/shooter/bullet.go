package shooter

import (
	"github.com/veandco/go-sdl2/sdl"
	"manveer/exp/components/common"
	"math"
)

type Bullet struct {
	common.Actor
}

var bulletCount = 0

func NewBullet(renderer *sdl.Renderer) *common.Actor {
	b := &common.Actor{}
	b.DrawWidth = Configs.BulletWidth
	b.DrawHeight = Configs.BulletHeight
	sr := common.NewSpriteRenderer(b, renderer, "sprites/player_bullet.bmp")
	b.AddComponent(sr)
	b.AddComponent(newBulletTrajectory(b, Configs.BulletSpeed))
	b.Active = false
	b.Tag = "bullet"
	col := common.Circle{
		Center: b.Position,
		Radius: Configs.BulletWidth,
	}
	b.Collisions = append(b.Collisions, col)
	return b
}

var BulletPool []*common.Actor

func InitBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullet := NewBullet(renderer)
		BulletPool = append(BulletPool, bullet)
		common.Actors = append(common.Actors, bullet)
	}
}

func bulletFromPool() (*common.Actor, bool) {
	for _, bul := range BulletPool {
		if !bul.Active {
			return bul, true
		}
	}
	return nil, false
}

type bulletTrajectory struct {
	container *common.Actor
	speed     float64
}

func newBulletTrajectory(container *common.Actor, speed float64) *bulletTrajectory {
	return &bulletTrajectory{
		container: container,
		speed:     speed,
	}
}

func (b *bulletTrajectory) OnUpdate() error {
	container := b.container
	container.Position.Y -= Configs.BulletSpeed * math.Cos(container.Angle*(math.Pi/180)) * common.Delta
	container.Position.X += Configs.BulletSpeed * math.Sin(container.Angle*(math.Pi/180)) * common.Delta
	if container.Position.Y < 0 {
		container.Active = false
	}
	b.container.Collisions[0].Center = container.Position

	return nil
}

func (b *bulletTrajectory) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (b *bulletTrajectory) OnCollision(other *common.Actor) error {
	b.container.Active = false
	return nil
}
