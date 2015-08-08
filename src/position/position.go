package position

import "math"

type Position struct {
	X float64
	Y float64
	Z float64
}

func NewPosition(x float64, y float64, z float64) *Position {
	position := Position{}
	position.X = x
	position.Y = y
	position.Z = z
	return &position
}

func (this Position) Distance(there Position) float64 {
	xNom := (this.X - (there.X)) * (this.X - (there.X))
	yNom := (this.Y - (there.Y)) * (this.Y - (there.Y))
	zNom := (this.Z - (there.Z)) * (this.Z - (there.Z))

	dist := math.Sqrt(xNom + yNom + zNom)
	return dist
}

func (this Position) Nothing() {

}
