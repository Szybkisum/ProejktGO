package main

import (
	"image/color"
	"math/rand"
)

type Grass struct {
	GameEntity
	ReproductionCooldown,
	MaxCooldown int
	isDead bool
}

func NewGrass(pos Position) *Grass {
	return &Grass{
		GameEntity: GameEntity{
			Pos:    pos,
			Color:  color.RGBA{R: 0, G: 255, B: 0, A: 255},
			Radius: 1.5,
		},
		ReproductionCooldown: 500,
		MaxCooldown:          1000,
	}
}

func (gr *Grass) Die() {
	gr.isDead = true
}

func (gr *Grass) IsDead() bool {
	return gr.isDead
}

func (gr *Grass) IsReadyToReproduce() bool {
	return gr.ReproductionCooldown <= 0
}

func (gr *Grass) RecoverFromReproduction() {
	if gr.ReproductionCooldown > 0 {
		gr.ReproductionCooldown--
	}
}

func (gr *Grass) StartReproductionCooldown() {
	gr.ReproductionCooldown = gr.MaxCooldown
}

func (gr *Grass) Reproduce() *Grass {
	offsetX := (rand.Float64()*2 - 1) * 10.0
	offsetY := (rand.Float64()*2 - 1) * 10.0
	newPos := Position{X: gr.Pos.X + offsetX, Y: gr.Pos.Y + offsetY}

	gr.StartReproductionCooldown()
	return NewGrass(newPos)
}

func (gr *Grass) Update(w *World) (newEntity Entity) {
	gr.RecoverFromReproduction()
	if gr.IsReadyToReproduce() {
		newEntity = gr.Reproduce()
		w.WorldBoundary.FitIntoBoundary(newEntity.GetPosition())
	}
	return
}
