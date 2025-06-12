package main

type Boundary struct {
	X, Y, Width, Height float64
}

func NewBoundary(X, Y, Width, Height float64) *Boundary {
	return &Boundary{
		X:X,
		Y:Y,
		Width: Width,
		Height: Height,
	}
}

func (b *Boundary) Contains(pos *Position) bool {
	halfWidth := b.Width / 2
	halfHeight := b.Height / 2
	return 
		pos.X >= b.X - halfWidth &&
		pos.X <= b.X + halfWidth &&
		pos.Y >= b.Y - halfHeight &&
		pos.Y <= b.Y + halfHeight
}


func (b *Boundary) GetEdges() (float64, float64, float64, float64) {
	halfWidth := b.Width / 2
	halfHeight := b.Height / 2

	left := b.X - halfWidth
	right := b.X + halfWidth
	up := b.Y - halfHeight
	down := b.Y + halfHeight

	return left, right, up, down
}

func (b *Boundary) FitIntoBoundary(pos *Position) {
	left, right, up, down := b.GetEdges()

	if left > pos.X {pos.X = left}
	if right < pos.X {pos.X = right}
	if up > pos.Y {pos.Y = up}
	if down < pos.Y {pos.Y = down}
}

func (b *Boundary) Intersects(other *Boundary) bool {
	left, right, up, down := b.GetEdges()
	otherLeft, otherRight, otherUp, otherDown := other.GetEdges()
	return !(
		left > otherRight || 
		right < otherLeft || 
		up > otherDown || 
		down < otherUp)
}