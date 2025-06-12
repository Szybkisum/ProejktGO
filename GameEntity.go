package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameEntity struct {
	Pos Position
	Color color.Color
	Radius float32
}

func (e *GameEntity) GetPosition() *Position {
	return &e.Pos
}

func (e *GameEntity) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(
		screen,
		float32(e.Pos.X),
		float32(e.Pos.Y),
		e.Radius,
		e.Color, 
		false,
	)
} 