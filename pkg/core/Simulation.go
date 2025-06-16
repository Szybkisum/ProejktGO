package core

import (
	"ProjektGO/pkg/config"
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	HISTORY_RECORD_INTERVAL = 15
	MAX_HISTORY_LENGTH      = 450
)

type Game struct {
	ScreenWidth, ScreenHeight int
	Capacity                  int
	World                     *World
	IsInitialized, IsPaused   bool
	historyTimer              int
}

func (g *Game) DrawPause(screen *ebiten.Image) {
	overlayColor := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	vector.DrawFilledRect(screen, 0, 0, float32(g.ScreenWidth), float32(g.ScreenHeight), overlayColor, false)
	pauseText := "PAUZA"
	textWidth := len(pauseText) * 7
	x := (g.ScreenWidth - textWidth) / 2
	y := g.ScreenHeight / 2
	text.Draw(screen, pauseText, basicfont.Face7x13, x, y, color.White)
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.IsPaused = !g.IsPaused
	}

	if g.IsPaused {
		return nil
	}

	if g.IsInitialized {
		var newlyBorn []Entity

		for _, entity := range g.World.GetAllEntities() {
			offspring := entity.Update(g.World)
			if offspring != nil {
				newlyBorn = append(newlyBorn, offspring)
			}
		}

		if g.World.IsGrassReadyToSpawn() {
			newlyBorn = append(newlyBorn, g.World.SpawnGrass()...)
		} else {
			g.World.GrassSpawnCooldown--
		}

		g.World.RemoveDeadEntities()

		g.World.AddNewEntities(newlyBorn)

		g.historyTimer++
		if g.historyTimer >= HISTORY_RECORD_INTERVAL {
			g.historyTimer = 0
			g.World.RabbitHistory = append(g.World.RabbitHistory, len(g.World.Rabbits))
			g.World.FoxHistory = append(g.World.FoxHistory, len(g.World.Foxes))
			g.World.GrassHistory = append(g.World.GrassHistory, len(g.World.Grass))
			if len(g.World.RabbitHistory) > MAX_HISTORY_LENGTH {
				g.World.RabbitHistory = g.World.RabbitHistory[1:]
				g.World.FoxHistory = g.World.FoxHistory[1:]
				g.World.GrassHistory = g.World.GrassHistory[1:]
			}
		}
	} else {
		g.IsInitialized = true
	}

	g.World.Quadtree = NewQuadTree(g.World.WorldBoundary, g.Capacity, 0)
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
		"FPS: %.2f\nKroliki: %d\nLisy: %d\nTrawa: %d\nSPACJA-PAUZA",
		fps, rabbitCount, foxCount, grassCount,
	)
	ebitenutil.DebugPrint(screen, msg)

	graphHeight := 100
	graphY := g.ScreenHeight - graphHeight

	bgColor := color.RGBA{R: 10, G: 10, B: 20, A: 210}
	vector.DrawFilledRect(screen, 0, float32(graphY), float32(g.ScreenWidth), float32(graphHeight), bgColor, false)

	maxPop := 1
	histories := [][]int{g.World.RabbitHistory, g.World.FoxHistory, g.World.GrassHistory}
	for _, history := range histories {
		for _, val := range history {
			if val > maxPop {
				maxPop = val
			}
		}
	}

	drawLine := func(history []int, lineColor color.Color) {
		if len(history) < 2 {
			return
		}
		for i := 1; i < len(history); i++ {
			x0 := float32(i-1) * float32(g.ScreenWidth) / float32(MAX_HISTORY_LENGTH-1)
			x1 := float32(i) * float32(g.ScreenWidth) / float32(MAX_HISTORY_LENGTH-1)
			y0 := float32(graphY + graphHeight - (history[i-1] * graphHeight / maxPop))
			y1 := float32(graphY + graphHeight - (history[i] * graphHeight / maxPop))

			vector.StrokeLine(screen, x0, y0, x1, y1, 1.5, lineColor, true)
		}
	}
	drawLine(g.World.GrassHistory, color.RGBA{R: 100, G: 255, B: 100, A: 255})
	drawLine(g.World.RabbitHistory, color.White)
	drawLine(g.World.FoxHistory, color.RGBA{R: 255, G: 165, B: 0, A: 255})

	if g.IsPaused {
		g.DrawPause(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func RunSimulation(cfg *config.SimulationConfig) {
	screenWidth := 800
	screenHeight := 800

	f64ScreenWidth := float64(screenWidth)
	f64ScreenHeight := float64(screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Lisy i Króliki")

	initialGrass := []*Grass{}
	for range cfg.InitialGrass {
		initialGrass = append(initialGrass, NewGrass(&Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}))
	}

	initialRabbits := []*Rabbit{}
	for range cfg.InitialRabbits {
		initialRabbits = append(initialRabbits, NewRabbit(&Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}, &cfg.RabbitParams))
	}

	initialFoxes := []*Fox{}
	for range cfg.InitialFoxes {
		initialFoxes = append(initialFoxes, NewFox(&Position{
			X: rand.Float64() * f64ScreenWidth,
			Y: rand.Float64() * f64ScreenHeight,
		}, &cfg.FoxParams))
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
		World: &World{
			Grass:              initialGrass,
			Rabbits:            initialRabbits,
			Foxes:              initialFoxes,
			WorldBoundary:      worldBoundary,
			Config:             cfg,
			GrassSpawnCooldown: cfg.GrassParams.GrassSpawnInterval,
		},
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
