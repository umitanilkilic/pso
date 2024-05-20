package pso

type Position struct {
	coordinates   []float64
	dimensionSize int
}

func (p *Position) UpdateCoordinate(coordinates ...float64) *Position {
	if len(coordinates) > p.dimensionSize {
		panic("Too many coordinates provided")
	}
	p.coordinates = coordinates
	return p
}

func (p *Position) GetCoordinates() []float64 {
	return p.coordinates
}

func (p *Position) GetDimension() int {
	return p.dimensionSize
}

func NewPosition(coords ...float64) Position {
	dim := len(coords)
	return Position{coordinates: coords, dimensionSize: dim}
}

func (p Position) Subtract(p2 Position) Position {
	if p.dimensionSize != p2.dimensionSize {
		panic("Dimensions do not match")
	}
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.coordinates[i] - p2.coordinates[i]
	}
	return Position{coordinates: result, dimensionSize: p.dimensionSize}
}

func (p Position) Add(p2 Position) Position {
	if p.dimensionSize != p2.dimensionSize {
		panic("Dimensions do not match")
	}
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.coordinates[i] + p2.coordinates[i]
	}
	return Position{coordinates: result, dimensionSize: p.dimensionSize}
}

func (p Position) Multiply(scalar float64) Position {
	result := make([]float64, p.dimensionSize)
	for i := 0; i < p.dimensionSize; i++ {
		result[i] = p.coordinates[i] * scalar
	}
	return Position{coordinates: result, dimensionSize: p.dimensionSize}
}
