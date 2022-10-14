package shooter

import "time"

type GameConfig struct {
	ScreenHeight         float64
	ScreenWidth          float64
	TargetTicksPerSecond float64

	PlayerBaseSpeed     float64
	PlayerSize          float64
	PlayerShotCoolDown  time.Duration
	PlayerRotationSpeed float64

	BasicEnemySize float64

	BulletHeight float64
	BulletWidth  float64
	BulletSpeed  float64
}

var Delta float64

var Configs = GameConfig{
	ScreenHeight:         960,
	ScreenWidth:          720,
	TargetTicksPerSecond: 60,

	PlayerBaseSpeed:     5,
	PlayerSize:          150,
	PlayerShotCoolDown:  time.Millisecond * 250,
	PlayerRotationSpeed: 1,

	BasicEnemySize: 70,

	BulletWidth:  10,
	BulletHeight: 50,
	BulletSpeed:  10,
}
