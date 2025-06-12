package main

type World struct {
	Quadtree *QuadTree
	Rabbits []*Rabbit
	Foxes []*Fox
	Grass []*Grass
    WorldBoundary *Boundary
}


func (w *World) GetAllEntities() []Entity {
    all := make([]Entity, 0, len(w.Rabbits) + len(w.Foxes) + len(w.Grass))
    for _, r := range w.Rabbits { all = append(all, r) }
    for _, f := range w.Foxes { all = append(all, f) }
    for _, gr := range w.Grass { all = append(all, gr) }
    return all
}

func (w *World) RemoveDeadEntities() {
    newRabbits := []*Rabbit{}
    for _, r := range w.Rabbits {
        if (!r.IsDead) {
            newRabbits = append(newRabbits, r)
        }
    }

    w.Rabbits = newRabbits
    newFoxes := []*Fox{}

    for _, f := range w.Foxes {
        if (!f.IsDead) {
            newFoxes = append(newFoxes, f)
        }
    }
    w.Foxes = newFoxes
} 