package pso

type Position struct {
	Coordinates   []float64
	dimensionSize int
}

func (p *Position) UpdateCoordinate(coordinates ...float64) *Position {
	if len(coordinates) > p.dimensionSize {
		panic("Too many coordinates provided")
	}
	p.Coordinates = coordinates
	return p
}

func (p *Position) GetCoordinates() []float64 {
	return p.Coordinates
}

func (p *Position) GetDimension() int {
	return p.dimensionSize
}

func NewPosition(coordinates ...float64) Position {
	dim := len(coordinates)
	return Position{Coordinates: coordinates, dimensionSize: dim}
}

func (p Position) Subtract(p2 Position) Position {
	if p.dimensionSize != p2.dimensionSize {
		panic("Dimensions do not match")
	}
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.Coordinates[i] - p2.Coordinates[i]
	}
	return Position{Coordinates: result, dimensionSize: p.dimensionSize}
}

func (p Position) Add(p2 Position) Position {
	if p.dimensionSize != p2.dimensionSize {
		panic("Dimensions do not match")
	}
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.Coordinates[i] + p2.Coordinates[i]
	}
	return Position{Coordinates: result, dimensionSize: p.dimensionSize}
}

func (p Position) Multiply(scalar float64) Position {
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.Coordinates[i] * scalar
	}
	return Position{Coordinates: result, dimensionSize: p.dimensionSize}
}
