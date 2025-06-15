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

func (r *Rabbit) IsDangerouslyClose(other Entity) bool {
	pos := r.GetPosition()
	otherPos := other.GetPosition()

	distSq := pos.CalculateDistanceSquared(otherPos)
	halfSeeingRange := r.SeeingRange / 2
	return distSq < halfSeeingRange*halfSeeingRange
}

func (r *Rabbit) RandomMove() {
	deltaX := ((rand.Float64() * 2) - 1) * r.Speed
	deltaY := ((rand.Float64() * 2) - 1) * r.Speed
	r.Pos.Move(deltaX, deltaY)
}

func (r *Rabbit) Eat(gr *Grass) {
	r.RecoverEnergy()
	gr.Die()
}

func (r *Rabbit) Reproduce(other *Rabbit) *Rabbit {
	midpointX := (r.Pos.X + other.Pos.X) / 2
	midpointY := (r.Pos.Y + other.Pos.Y) / 2

	offsetX := (rand.Float64()*2 - 1) * 2.0
	offsetY := (rand.Float64()*2 - 1) * 2.0

	newPos := Position{X: midpointX + offsetX, Y: midpointY + offsetY}
	newRabbit := NewRabbit(newPos)

	r.StartReproductionCooldown()
	other.StartReproductionCooldown()

	return newRabbit
}

func (r *Rabbit) MoveToward(other Entity) {
	pos := r.GetPosition()
	targetPos := other.GetPosition()

	directionX := targetPos.X - pos.X
	directionY := targetPos.Y - pos.Y

	length := math.Sqrt(directionX*directionX + directionY*directionY)
	if length > 0 {
		moveX := (directionX / length) * r.Speed
		moveY := (directionY / length) * r.Speed
		pos.Move(moveX, moveY)
	}
}

func (r *Rabbit) MoveAwayFrom(other Entity) {
	pos := r.GetPosition()
	otherPos := other.GetPosition()

	directionX := pos.X - otherPos.X
	directionY := pos.Y - otherPos.Y

	length := math.Sqrt(directionX*directionX + directionY*directionY)
	if length > 0 {
		moveX := (directionX / length) * r.Speed
		moveY := (directionY / length) * r.Speed
		pos.Move(moveX, moveY)
	}
}

func (r *Rabbit) Update(w *World) (newEntity Entity) {
	r.Metabolise()

	pos := r.GetPosition()
	doubledSeeingRange := r.SeeingRange * 2
	searchArea := Boundary{X: pos.X, Y: pos.Y, Width: doubledSeeingRange, Height: doubledSeeingRange}
	found := w.Quadtree.Query(&searchArea)

	if r.IsHungry() {
		fox, grass := getClosestFoxAndGrass(r, found)
		if fox != nil && r.IsDangerouslyClose(fox) {
			r.MoveAwayFrom(fox)
		} else if grass != nil {
			if r.IsInRange(grass) {
				r.Eat(grass)
			} else {
				r.MoveToward(grass)
			}
		} else {
			r.RandomMove()
		}
	} else if r.IsReadyToReproduce() {
		fox, rabbit := getClosestFoxAndRabbitReadyToReproduce(r, found)
		if fox != nil && r.IsDangerouslyClose(fox) {
			r.MoveAwayFrom(fox)
		}
		if rabbit != nil {
			if r.IsInRange(rabbit) {
				newEntity = r.Reproduce(rabbit)
				w.WorldBoundary.FitIntoBoundary(newEntity.GetPosition())
			} else {
				r.MoveToward(rabbit)
			}
		} else {
			r.RandomMove()
		}
	} else {
		fox := getClosestFox(r, found)
		if fox != nil && r.IsDangerouslyClose(fox) {
			r.MoveAwayFrom(fox)
		} else {
			r.RandomMove()
		}
	}
	w.WorldBoundary.FitIntoBoundary(&r.Pos)
	return
}
