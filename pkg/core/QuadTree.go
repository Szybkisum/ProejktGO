package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const MAX_DEPTH = 100

type QuadTree struct {
	Boundary  *Boundary
	Capacity  int
	Depth     int
	Entities  []Entity
	NorthWest *QuadTree
	NorthEast *QuadTree
	SouthWest *QuadTree
	SouthEast *QuadTree
	IsDivided bool
}

func NewQuadTree(Boundary *Boundary, capacity int, depth int) *QuadTree {
	return &QuadTree{
		Boundary:  Boundary,
		Capacity:  capacity,
		Depth:     depth,
		IsDivided: false,
	}
}

func (qt *QuadTree) Subdivide() {
	newWidth := qt.Boundary.Width / 2
	newHeight := qt.Boundary.Height / 2
	halfNewWidth := newWidth / 2
	halfNewHeight := newHeight / 2
	nextDepth := qt.Depth + 1
	qt.NorthWest = NewQuadTree(NewBoundary(qt.Boundary.X-halfNewWidth, qt.Boundary.Y-halfNewHeight, newWidth, newHeight), qt.Capacity, nextDepth)
	qt.NorthEast = NewQuadTree(NewBoundary(qt.Boundary.X+halfNewWidth, qt.Boundary.Y-halfNewHeight, newWidth, newHeight), qt.Capacity, nextDepth)
	qt.SouthWest = NewQuadTree(NewBoundary(qt.Boundary.X-halfNewWidth, qt.Boundary.Y+halfNewHeight, newWidth, newHeight), qt.Capacity, nextDepth)
	qt.SouthEast = NewQuadTree(NewBoundary(qt.Boundary.X+halfNewWidth, qt.Boundary.Y+halfNewHeight, newWidth, newHeight), qt.Capacity, nextDepth)
	qt.IsDivided = true

	for _, e := range qt.Entities {
		qt.Insert(e)
	}

	qt.Entities = nil
}

func (qt *QuadTree) Insert(e Entity) {
	if !qt.Boundary.Contains(e.GetPosition()) {
		return
	} else if qt.Depth >= MAX_DEPTH {
		qt.Entities = append(qt.Entities, e)
		return
	} else if !qt.IsDivided {
		if len(qt.Entities) < qt.Capacity {
			qt.Entities = append(qt.Entities, e)
			return
		} else {
			qt.Subdivide()
		}
	}
	pos := e.GetPosition()
	if qt.NorthWest.Boundary.Contains(pos) {
		qt.NorthWest.Insert(e)
	} else if qt.NorthEast.Boundary.Contains(pos) {
		qt.NorthEast.Insert(e)
	} else if qt.SouthWest.Boundary.Contains(pos) {
		qt.SouthWest.Insert(e)
	} else if qt.SouthEast.Boundary.Contains(pos) {
		qt.SouthEast.Insert(e)
	}
}

func (qt *QuadTree) Draw(screen *ebiten.Image) {
	halfW := qt.Boundary.Width / 2
	halfH := qt.Boundary.Height / 2
	x1 := float32(qt.Boundary.X - halfW)
	x2 := float32(qt.Boundary.X + halfW)
	y1 := float32(qt.Boundary.Y - halfH)
	y2 := float32(qt.Boundary.Y + halfH)
	lineColor := color.RGBA{R: 50, G: 50, B: 50, A: 255}
	strokeWidth := float32(1)
	vector.StrokeLine(screen, x1, y1, x2, y1, strokeWidth, lineColor, false)
	vector.StrokeLine(screen, x1, y2, x2, y2, strokeWidth, lineColor, false)
	vector.StrokeLine(screen, x1, y1, x1, y2, strokeWidth, lineColor, false)
	vector.StrokeLine(screen, x2, y1, x2, y2, strokeWidth, lineColor, false)
	if qt.IsDivided {
		qt.NorthWest.Draw(screen)
		qt.NorthEast.Draw(screen)
		qt.SouthWest.Draw(screen)
		qt.SouthEast.Draw(screen)
	}
}

func (qt *QuadTree) Query(searchArea *Boundary) []Entity {
	found := []Entity{}
	if !qt.Boundary.Intersects(searchArea) {
		return found
	}
	if qt.IsDivided {
		found = append(found, qt.NorthWest.Query(searchArea)...)
		found = append(found, qt.NorthEast.Query(searchArea)...)
		found = append(found, qt.SouthWest.Query(searchArea)...)
		found = append(found, qt.SouthEast.Query(searchArea)...)
	} else {
		for _, e := range qt.Entities {
			if !e.IsDead() && searchArea.Contains(e.GetPosition()) {
				found = append(found, e)
			}
		}
	}
	return found
}
