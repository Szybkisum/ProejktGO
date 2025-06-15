package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameEntity struct {
	Pos    Position
	Color  color.Color
	Radius float32
}

func (e *GameEntity) GetPosition() *Position {
	return &e.Pos
}

func (e *GameEntity) GetRadius() float64 {
	return float64(e.Radius)
}

func (e *GameEntity) IsInRange(other Entity) bool {
	targetPos := other.GetPosition()
	interactionDistance := e.GetRadius() + other.GetRadius()

	distSq := e.Pos.CalculateDistanceSquared(targetPos)
	return distSq <= interactionDistance*interactionDistance
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
