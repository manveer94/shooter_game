package common

import (
	"math"
)

type Circle struct {
	Center Vector
	Radius float64
}

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) + math.Pow(c2.Center.Y-c1.Center.Y, 2))
	return dist <= c1.Radius+c2.Radius
}

func CheckCollisions() error {
	for i := 0; i < len(Actors)-1; i++ {
		for j := i + 1; j < len(Actors); j++ {
			for _, c1 := range Actors[i].Collisions {
				for _, c2 := range Actors[j].Collisions {
					if Actors[i].Active && Actors[j].Active && collides(c1, c2) {
						err := Actors[i].Collision(Actors[j])
						if err != nil {
							return err
						}
						err = Actors[j].Collision(Actors[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
