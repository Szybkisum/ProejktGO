package main

import (
	"image/color"
	"math/rand"
)

type Rabbit struct {
	LivingEntity
}

func NewRabbit(pos *Position) *Rabbit {
	return &Rabbit{
		LivingEntity: LivingEntity{
			GameEntity: GameEntity{
				Pos:    pos,
				Color:  color.RGBA{255, 255, 255, 255},
				Radius: 1.5,
				isDead: false,
			},
			Speed:                1.5,
			SeeingRange:          50.0,
			Energy:               500.0,
			MaxEnergy:            1000.0,
			EnergyToReproduce:    700.0,
			ReproductionCooldown: 0,
			MaxCooldown:          150,
		},
	}
}

func (r *Rabbit) Eat(gr *Grass) {
	r.RecoverEnergy()
	gr.Die()
}

func (r *Rabbit) TargetFood(gr *Grass) {
	if r.IsInRange(gr) {
		r.Eat(gr)
	} else {
		r.MoveToward(gr)
	}
}

func (r *Rabbit) TargetPartner(other *Rabbit) (newEntity Entity) {
	if r.IsInRange(other) {
		newEntity = r.Reproduce(other)
	} else {
		r.MoveToward(other)
	}
	return
}

func (r *Rabbit) Reproduce(other *Rabbit) (newEntity Entity) {
	otherPos := other.GetPosition()

	midpointX := (r.Pos.X + otherPos.X) / 2
	midpointY := (r.Pos.Y + otherPos.Y) / 2

	offsetX := (rand.Float64()*2 - 1) * 2.0
	offsetY := (rand.Float64()*2 - 1) * 2.0

	newPos := &Position{X: midpointX + offsetX, Y: midpointY + offsetY}
	newEntity = NewRabbit(newPos)

	r.StartReproductionCooldown()
	other.StartReproductionCooldown()

	return newEntity
}

func (r *Rabbit) Update(w *World) (newEntity Entity) {
	r.Metabolise()

	pos := r.GetPosition()
	doubledSeeingRange := r.SeeingRange * 2
	searchArea := Boundary{X: pos.X, Y: pos.Y, Width: doubledSeeingRange, Height: doubledSeeingRange}
	found := w.Quadtree.Query(&searchArea)

	rabbit, fox, grass := filterRabbitInterests(r, found)

	if fox != nil && r.IsDangerouslyClose(fox) {
		r.MoveAwayFrom(fox)
	} else if r.IsHungry() && grass != nil {
		r.TargetFood(grass)
	} else if r.IsReadyToReproduce() && rabbit != nil {
		if newEntity = r.TargetPartner(rabbit); newEntity != nil {
			w.WorldBoundary.FitIntoBoundary(newEntity.GetPosition())
		}
	} else {
		r.RandomMove()
	}
	w.WorldBoundary.FitIntoBoundary(r.Pos)
	return
}
