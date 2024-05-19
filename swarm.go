package pso

import "math/rand"

type FitnessFunction func(Position) float64

type Swarm struct {
	Inertia     float64
	ConstantOne float64
	ConstantTwo float64
	GlobalBest  Position
	Particles   []*Particle
	FitnessFunc FitnessFunction
}

func NewSwarm(inertia, c1, c2 float64, GlobalBest Position, particles []*Particle) *Swarm {
	return &Swarm{Inertia: inertia, ConstantOne: c1, ConstantTwo: c2, Particles: particles, GlobalBest: GlobalBest}
}

func (s *Swarm) CalculateVelocity() {
	for _, particle := range s.Particles {
		// Calculate the velocity
		// v = w*v + c1*r1*(p-x) + c2*r2*(g-x)
		inertiaPart := particle.GetVelocity().Multiply(s.Inertia)
		c1Part := particle.GetBestPosition().Subtract(*particle.GetPosition()).Multiply(s.ConstantOne * rand.Float64())
		c2Part := s.GlobalBest.Subtract(*particle.GetPosition()).Multiply(s.ConstantTwo * rand.Float64())
		newVelocity := inertiaPart.Add(c1Part).Add(c2Part)
		particle.UpdateVelocity(newVelocity)

	}
}

func (s *Swarm) UpdatePosition() {
	for _, particle := range s.Particles {
		newPosition := particle.GetPosition().Add(*particle.GetVelocity())
		particle.UpdatePosition(newPosition)
	}
}

func (s *Swarm) UpdateBestPosition() {
	for _, particle := range s.Particles {
		// Update the best position
		// if f(x) < f(p) then p = x
		if s.FitnessFunc(*particle.GetPosition()) < s.FitnessFunc(*particle.GetBestPosition()) {
			particle.BestPosition = particle.GetPosition()
		}
	}
}

func (s *Swarm) UpdateGlobalBest() {
	// Update the global best position
	// if f(x) < f(g) then g = x
	for _, particle := range s.Particles {
		if s.FitnessFunc(*particle.GetPosition()) < s.FitnessFunc(s.GlobalBest) {
			s.GlobalBest = *particle.GetPosition()
		}
	}
}

func (s *Swarm) Solve(iterations int, epsilon float64) {
	s.UpdateBestPosition()
	for i := 0; i < iterations; i++ {
		s.CalculateVelocity()
		s.UpdatePosition()
		s.UpdateBestPosition()
		s.UpdateGlobalBest()
	}
}
