package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
    GetPosition() *Position
    Update(world *World)
	Draw(screen *ebiten.Image)
} 