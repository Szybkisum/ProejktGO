package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	ScreenWidth, ScreenHeight int
	Capacity                  int
	World                     *World
	IsInitialized, IsPaused   bool
}

const GLOBAL_GRASS_SPAWN_INTERVAL = 120

func (g *Game) Update() error {

	if g.IsInitialized {
		var newlyBorn []Entity

		for _, entity := range g.World.GetAllEntities() {
			offspring := entity.Update(g.World)
			if offspring != nil {
				newlyBorn = append(newlyBorn, offspring)
			}
		}

		if g.World.IsGlobalGrassReadyToSpawn() {
			newlyBorn = append(newlyBorn, g.World.SpawnGlobalGrass())
		} else {
			g.World.GlobalGrassSpawnCooldown--
		}

		g.World.RemoveDeadEntities()

		g.World.AddNewEntities(newlyBorn)
	} else {
		g.IsInitialized = true
	}

	g.World.Quadtree = NewQuadTree(g.World.WorldBoundary, g.Capacity)
	for _, entity := range g.World.GetAllEntities() {
		g.World.Quadtree.Insert(entity)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Quadtree.Draw(screen)
	for _, entity := range g.World.GetAllEntities() {
		entity.Draw(screen)
	}

	fps := ebiten.ActualFPS()
	rabbitCount := len(g.World.Rabbits)
	foxCount := len(g.World.Foxes)
	grassCount := len(g.World.Grass)
	msg := fmt.Sprintf(
		"FPS: %.2f\nKroliki: %d\nLisy: %d\nTrawa: %d",
		fps, rabbitCount, foxCount, grassCount,
	)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	screenWidth := 900
	screenHeight := 900

	f64ScreenWidth := float64(screenWidth)
	f64ScreenHeight := float64(screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Lisy i Króliki")

	initialGrass := []*Grass{}
	for i := 0; i < 1000; i++ {
		initialGrass = append(initialGrass, NewGrass(Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}))
	}

	initialRabbits := []*Rabbit{}
	for i := 0; i < 3000; i++ {
		initialRabbits = append(initialRabbits, NewRabbit(Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}))
	}

	initialFoxes := []*Fox{}
	for i := 0; i < 200; i++ {
		initialFoxes = append(initialFoxes, NewFox(Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}))
	}
	worldBoundary := &Boundary{
		X:      f64ScreenWidth / 2,
		Y:      f64ScreenHeight / 2,
		Width:  f64ScreenWidth,
		Height: f64ScreenHeight,
	}

	g := &Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Capacity:     4,
		World:        &World{Grass: initialGrass, Rabbits: initialRabbits, Foxes: initialFoxes, WorldBoundary: worldBoundary, GlobalGrassSpawnCooldown: GLOBAL_GRASS_SPAWN_INTERVAL},
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
