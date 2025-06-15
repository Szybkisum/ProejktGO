package main

import (
	"image/color"
	"math"
	"math/rand"
)

type Fox struct {
	LivingEntity
	GameEntity
}

func NewFox(pos Position) *Fox {
	return &Fox{
		LivingEntity: LivingEntity{
			Speed:                2,
			SeeingRange:          75.0,
			Energy:               750.0,
			MaxEnergy:            1500.0,
			EnergyToReproduce:    1000.0,
			ReproductionCooldown: 0,
			MaxCooldown:          600,
		},
		GameEntity: GameEntity{
			Pos:    pos,
			Color:  color.RGBA{255, 165, 0, 255},
			Radius: 1.5,
		},
	}
}

func (f *Fox) RandomMove() {
	deltaX := ((rand.Float64() * 2) - 1) * f.Speed
	deltaY := ((rand.Float64() * 2) - 1) * f.Speed
	f.Pos.Move(deltaX, deltaY)
}

func (f *Fox) Eat(r *Rabbit) {
	f.RecoverEnergy()
	r.Die()
}

func (f *Fox) Reproduce(other *Fox) *Fox {
	midpointX := (f.Pos.X + other.Pos.X) / 2
	midpointY := (f.Pos.Y + other.Pos.Y) / 2

	offsetX := (rand.Float64()*2 - 1) * 2.0
	offsetY := (rand.Float64()*2 - 1) * 2.0

	newPos := Position{X: midpointX + offsetX, Y: midpointY + offsetY}
	newFox := NewFox(newPos)

	f.StartReproductionCooldown()
	other.StartReproductionCooldown()

	return newFox
}

func (f *Fox) MoveToward(other Entity) {
	pos := f.GetPosition()
	targetPos := other.GetPosition()

	directionX := targetPos.X - pos.X
	directionY := targetPos.Y - pos.Y

	length := math.Sqrt(directionX*directionX + directionY*directionY)
	if length > 0 {
		moveX := (directionX / length) * f.Speed
		moveY := (directionY / length) * f.Speed
		pos.Move(moveX, moveY)
	}
}

func (f *Fox) Update(w *World) (newEntity Entity) {
	f.Metabolise()

	pos := f.GetPosition()
	doubledSeeingRange := f.SeeingRange * 2
	searchArea := Boundary{X: pos.X, Y: pos.Y, Width: doubledSeeingRange, Height: doubledSeeingRange}
	found := w.Quadtree.Query(&searchArea)

	if f.IsHungry() {
		rabbit := getClosestRabbit(f, found)
		if rabbit != nil {
			if f.IsInRange(rabbit) {
				f.Eat(rabbit)
			} else {
				f.MoveToward(rabbit)
			}
		} else {
			f.RandomMove()
		}
	} else if f.IsReadyToReproduce() {
		fox := getClosestFoxReadyToReproduce(f, found)
		if fox != nil {
			if f.IsInRange(fox) {
				newEntity = f.Reproduce(fox)
				w.WorldBoundary.FitIntoBoundary(newEntity.GetPosition())
			} else {
				f.MoveToward(fox)
			}
		} else {
			f.RandomMove()
		}
	} else {
		f.RandomMove()
	}
	w.WorldBoundary.FitIntoBoundary(&f.Pos)
	return
}
