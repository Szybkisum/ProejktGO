package main

import (
	"image/color"
	"math"
	"math/rand"
)

type Rabbit struct {
	LivingEntity
	GameEntity
}

func NewRabbit(pos Position) *Rabbit {
	return &Rabbit{
		LivingEntity: LivingEntity{
			Speed:                1.5,
			SeeingRange:          50.0,
			Energy:               500.0,
			MaxEnergy:            1000.0,
			EnergyToReproduce:    700.0,
			ReproductionCooldown: 0,
			MaxCooldown:          300,
		},
		GameEntity: GameEntity{
			Pos:    pos,
			Color:  color.RGBA{255, 255, 255, 255},
			Radius: 1.5,
		},
	}
}

func (r *Rabbit) RandomMove() {
	deltaX := ((rand.Float64() * 2) - 1) * r.Speed
	deltaY := ((rand.Float64() * 2) - 1) * r.Speed
	r.Pos.Move(deltaX, deltaY)
}

func (r *Rabbit) MoveToward(other Entity) {
	pos := r.GetPosition()
	targetPos := other.GetPosition()
// dist := math.Sqrt(r.Pos.CalculateDistanceSquared(targetPos))
// if dist <= r.Speed {
// pos.MoveTo(targetPos.X, targetPos.Y)
// } else {
        directionX := targetPos.X - pos.X
        directionY := targetPos.Y - pos.Y
        length := math.Sqrt(directionX*directionX + directionY*directionY)
        if length > 0 {
            moveX := (directionX / length) * r.Speed
            moveY := (directionY / length) * r.Speed
            pos.Move(moveX, moveY)
        }
// }
}

func (r *Rabbit) Decide(w *World) {
	pos := r.GetPosition()
	doubledSeeingRange := r.SeeingRange * 2
	searchArea := Boundary{X: pos.X, Y: pos.Y, Width: doubledSeeingRange, Height: doubledSeeingRange}
	found := w.Quadtree.Query(&searchArea)

	if r.IsHungry() {
		_, grass := getClosestFoxAndGrass(r, found)
		if grass != nil {
			r.MoveToward(grass)
			return
		}
	// } else if r.IsReadyToReproduce() {
	// fox, rabbit := getClosestFoxAndRabbitReadyToReproduce(r, found)
	}
	r.RandomMove()
}

func (r *Rabbit) Update(w *World) {
	r.Metabolise()
	r.Decide(w)
	w.WorldBoundary.FitIntoBoundary(&r.Pos)
}

func getClosestFoxAndGrass(r *Rabbit, entities []Entity) (fox Entity, grass Entity) {
	closestFoxDistanceSq := math.MaxFloat64
	closestGrassDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, e := range entities {
		otherPos = e.GetPosition()
		dist = otherPos.CalculateDistanceSquared(r.GetPosition())
		switch e.(type) {
			case *Fox:
				if closestFoxDistanceSq >= dist {
					closestFoxDistanceSq = dist
					fox = e
			}

			case *Grass:
				if closestGrassDistanceSq >= dist {
					closestGrassDistanceSq = dist
					grass = e
			}

			default:
				continue
		}
	}
	return
}


// func getClosestFoxAndRabbitReadyToReproduce(r *Rabbit, entities []Entity) (fox Entity, rabbit Entity) {
// seeingDistanceSq := r.SeeingRange * r.SeeingRange
// closestFoxDistanceSq := seeingDistanceSq
// closestRabbitDistanceSq := seeingDistanceSq
// var otherPos *Position
// var dist float64
// for _, e := range entities {
// otherPos = e.GetPosition()
// dist = otherPos.CalculateDistanceSquared(r.GetPosition())
// switch e.(type) {
// case *Fox:
// if closestFoxDistanceSq >= dist {
// closestFoxDistanceSq = dist
// }
// case *Rabbit:
// // sprawdzi─ç IsReadyToReproduce
// if closestRabbitDistanceSq >= dist {
// closestRabbitDistanceSq = dist
// }
// default:
// continue
// }
// }
// return
// } 