package core

type Position struct {
	X, Y float64
}

func (pos *Position) Move(deltaX, deltaY float64) {
	pos.X += deltaX
	pos.Y += deltaY
}

func (pos *Position) MoveTo(x, y float64) {
	pos.X = x
	pos.Y = y
}

func (pos *Position) CalculateDistanceSquared(other *Position) float64 {
	distX := pos.X - other.X
	distY := pos.Y - other.Y

	return distX*distX + distY*distY
}
