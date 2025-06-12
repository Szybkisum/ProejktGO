package main

import (
	"image/color"
	"math/rand"
)

type Fox struct {
	LivingEntity
	GameEntity
}

func NewFox(pos Position) *Fox {
	return &Fox{
		LivingEntity: LivingEntity{
			Speed:              2,
			SeeingRange:        75.0,
			Energy:            750.0,
			MaxEnergy:          1500.0,
			EnergyToReproduce:  1000.0,
			ReproductionCooldown: 0,
			MaxCooldown:        600,
		},
		GameEntity: GameEntity{
			Pos: pos,
			Color: color.RGBA{255, 165, 0, 255},
			Radius: 1.5,
		},
	}
}

func (f *Fox) RandomMove() {
	deltaX := ((rand.Float64() * 2) - 1) * f.Speed
	deltaY := ((rand.Float64() * 2) - 1) * f.Speed
	f.Pos.Move(deltaX, deltaY)
}

func (f *Fox) Update(w *World) {
	f.Metabolise()
	f.RandomMove()
	w.WorldBoundary.FitIntoBoundary(&f.Pos)
} 