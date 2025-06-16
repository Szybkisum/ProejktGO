package core

import (
	"ProjektGO/pkg/config"
	"math/rand/v2"
)

type World struct {
	Quadtree           *QuadTree
	Rabbits            []*Rabbit
	Foxes              []*Fox
	Grass              []*Grass
	WorldBoundary      *Boundary
	Config             *config.SimulationConfig
	GrassSpawnCooldown int
}

func (w *World) IsGrassReadyToSpawn() bool {
	return w.GrassSpawnCooldown <= 0
}

func (w *World) StartGrassSpawnCooldown() {
	w.GrassSpawnCooldown = w.Config.GrassParams.GrassSpawnInterval
}

func (w *World) SpawnGrass() (newGrass []Entity) {
	for range w.Config.GrassParams.GrassSpawnCount {
		newGrass = append(newGrass, NewGrass(&Position{
			X: rand.Float64() * w.WorldBoundary.Width,
			Y: rand.Float64() * w.WorldBoundary.Height,
		}))
	}
	w.StartGrassSpawnCooldown()
	return
}

func (w *World) GetAllEntities() []Entity {
	all := make([]Entity, 0, len(w.Rabbits)+len(w.Foxes)+len(w.Grass))
	for _, r := range w.Rabbits {
		all = append(all, r)
	}
	for _, f := range w.Foxes {
		all = append(all, f)
	}
	for _, gr := range w.Grass {
		all = append(all, gr)
	}
	return all
}

func (w *World) RemoveDeadEntities() {
	newRabbits := []*Rabbit{}
	for _, r := range w.Rabbits {
		if !r.IsDead() {
			newRabbits = append(newRabbits, r)
		}
	}
	w.Rabbits = newRabbits

	newGrass := []*Grass{}
	for _, gr := range w.Grass {
		if !gr.IsDead() {
			newGrass = append(newGrass, gr)
		}
	}
	w.Grass = newGrass

	newFoxes := []*Fox{}
	for _, f := range w.Foxes {
		if !f.IsDead() {
			newFoxes = append(newFoxes, f)
		}
	}
	w.Foxes = newFoxes
}

func (w *World) AddNewEntities(entities []Entity) {
	for _, e := range entities {
		switch eType := e.(type) {
		case *Rabbit:
			w.Rabbits = append(w.Rabbits, eType)
		case *Fox:
			w.Foxes = append(w.Foxes, eType)
		case *Grass:
			w.Grass = append(w.Grass, eType)
		}
	}
}
