package main

import (
	"image/color"
	"math/rand"
)

type Fox struct {
	LivingEntity
}

func NewFox(pos *Position) *Fox {
	return &Fox{
		LivingEntity: LivingEntity{
			GameEntity: GameEntity{
				Pos:    pos,
				Color:  color.RGBA{255, 165, 0, 255},
				Radius: 1.5,
				isDead: false,
			},
			Speed:                2,
			SeeingRange:          75.0,
			Energy:               750.0,
			MaxEnergy:            1500.0,
			EnergyToReproduce:    1000.0,
			ReproductionCooldown: 0,
			MaxCooldown:          600,
		},
	}
}

func (f *Fox) Eat(r *Rabbit) {
	f.RecoverEnergy()
	r.Die()
}

func (f *Fox) TargetFood(r *Rabbit) {
	if f.IsInRange(r) {
		f.Eat(r)
	} else {
		f.MoveToward(r)
	}
}

func (f *Fox) TargetPartner(other *Fox) (newEntity Entity) {
	if f.IsInRange(other) {
		newEntity = f.Reproduce(other)
	} else {
		f.MoveToward(other)
	}
	return
}

func (f *Fox) Reproduce(other *Fox) (newEntity Entity) {
	otherPos := other.GetPosition()

	midpointX := (f.Pos.X + otherPos.X) / 2
	midpointY := (f.Pos.Y + otherPos.Y) / 2

	offsetX := (rand.Float64()*2 - 1) * 2.0
	offsetY := (rand.Float64()*2 - 1) * 2.0

	newPos := &Position{X: midpointX + offsetX, Y: midpointY + offsetY}
	newEntity = NewFox(newPos)

	f.StartReproductionCooldown()
	other.StartReproductionCooldown()

	return newEntity
}

func (f *Fox) Update(w *World) (newEntity Entity) {
	f.Metabolise()

	pos := f.GetPosition()
	doubledSeeingRange := f.SeeingRange * 2
	searchArea := Boundary{X: pos.X, Y: pos.Y, Width: doubledSeeingRange, Height: doubledSeeingRange}
	found := w.Quadtree.Query(&searchArea)

	rabbit, fox := filterFoxInterests(f, found)
	if f.IsHungry() && rabbit != nil {
		f.TargetFood(rabbit)
	} else if f.IsReadyToReproduce() && fox != nil {
		if newEntity = f.TargetPartner(fox); newEntity != nil {
			w.WorldBoundary.FitIntoBoundary(newEntity.GetPosition())
		}
	} else {
		f.RandomMove()
	}
	w.WorldBoundary.FitIntoBoundary(f.Pos)
	return
}
