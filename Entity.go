package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetPosition() *Position
	GetRadius() float64
	IsDead() bool
	Die()
	Update(world *World) Entity
	Draw(screen *ebiten.Image)
}
