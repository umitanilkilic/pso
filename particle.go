package pso

type Particle struct {
	ID           int
	Position     *Position
	Velocity     *Position
	BestPosition *Position
	BestFitness  *float64
}

func (p *Particle) UpdatePosition(position Position) *Particle {
	p.Position = &position
	return p
}

func (p *Particle) UpdateVelocity(velocity Position) *Particle {
	p.Velocity = &velocity
	return p
}

func (p *Particle) UpdateBestPosition(bestPosition Position) *Particle {
	p.BestPosition = &bestPosition
	return p
}

func (p *Particle) GetVelocity() *Position {
	return p.Velocity
}
func (p *Particle) GetPosition() *Position {
	return p.Position
}
func (p *Particle) GetBestPosition() *Position {
	return p.BestPosition
}

func NewParticle(id int, position Position, velocity Position) *Particle {
	return &Particle{
		ID:           id,
		Position:     &position,
		Velocity:     &velocity,
		BestPosition: &position,
	}
}
