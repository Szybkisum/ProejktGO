package main

import (
	"image/color"
)

type Grass struct {
	GameEntity
}

func NewGrass(pos *Position) *Grass {
	return &Grass{
		GameEntity: GameEntity{
			Pos:    pos,
			Color:  color.RGBA{R: 0, G: 255, B: 0, A: 255},
			Radius: 1.5,
			isDead: false,
		},
	}
}

func (gr *Grass) Update(w *World) (newEntity Entity) {
	return nil
}
