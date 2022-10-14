package shooter

import (
	"math"
)

type circle struct {
	center Vector
	radius float64
}

func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.X-c1.center.X, 2) + math.Pow(c2.center.Y-c1.center.Y, 2))
	return dist <= c1.radius+c2.radius
}

func CheckCollisions() error {
	for i := 0; i < len(Actors)-1; i++ {
		for j := i + 1; j < len(Actors); j++ {
			for _, c1 := range Actors[i].collisions {
				for _, c2 := range Actors[j].collisions {
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
