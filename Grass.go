package main

import (
	"image/color"
)

type Grass struct {
	GameEntity
	GrowthLevel, MaxGrowth float64
}

func NewGrass(pos Position) *Grass {
    return &Grass{
		GameEntity: GameEntity{
			Pos: pos,
			Color: color.RGBA{R: 0, G: 255, B: 0, A: 255},
			Radius: 1.5,
		},
        GrowthLevel: 1.0,
        MaxGrowth:   100.0,
    }
}

func (gr *Grass) Update(w *World) {} 