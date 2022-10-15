package levels

import (
	"manveer/exp/components/shooter"
)

func CreateShooterLevel() *Level {

	level := Level{
		Name:         "Shooter",
		onStart:      shooter.Initialize,
		windowWidth:  shooter.Configs.ScreenWidth,
		windowHeight: shooter.Configs.ScreenHeight,
	}
	return &level
}
