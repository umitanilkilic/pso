package pso

import (
	"math/rand"
)

type FitnessFunction func(Position) float64

type Swarm struct {
	// Inertia is the inertia of the swarm
	Inertia float64
	// ConstantOne is the social constant
	ConstantOne float64
	// ConstantTwo is the cognitive constant
	ConstantTwo float64
	// GlobalBest is the best position found by the swarm
	GlobalBest Position
	// Particles is a list of particles that the swarm will use to search the space
	Particles []*Particle
	// FitnessFunc is the function that the swarm will use to evaluate the fitness of a position
	FitnessFunc FitnessFunction
	// Boundaries is the boundaries of the search space
	BoundingRectangleSize float64
}

func (s *Swarm) CalculateVelocity() {
	for _, particle := range s.Particles {
		inertiaPart := particle.GetVelocity().Multiply(s.Inertia)
		c1Part := particle.GetBestPosition().Subtract(*particle.GetPosition()).Multiply(s.ConstantOne * rand.Float64() * 2)
		c2Part := s.GetGlobalBest().Subtract(*particle.GetPosition()).Multiply(s.ConstantTwo * rand.Float64() * 2)
		newVelocity := inertiaPart.Add(c1Part).Add(c2Part)
		particle.UpdateVelocity(newVelocity)
	}
}

func (s *Swarm) GetGlobalBest() *Position {
	return &s.GlobalBest
}

func (s *Swarm) UpdatePosition() {
	for _, particle := range s.Particles {
		newPosition := particle.GetPosition().Add(*particle.GetVelocity())
		s.limitWithinBoundaries(&newPosition)
		particle.UpdatePosition(newPosition)
	}
}

func (s *Swarm) UpdateBestPosition() {
	for _, particle := range s.Particles {
		if s.FitnessFunc(*particle.GetPosition()) < s.FitnessFunc(*particle.GetBestPosition()) {
			particle.BestPosition = particle.GetPosition()
		}
	}
}

func (s *Swarm) UpdateGlobalBest() {
	for _, particle := range s.Particles {
		if s.FitnessFunc(*particle.GetPosition()) < s.FitnessFunc(*s.GetGlobalBest()) {
			s.GlobalBest = *particle.GetPosition()
		}
	}
}

func (s *Swarm) limitWithinBoundaries(position *Position) {
	for i := 0; i < position.dimensionSize; i++ {
		if position.coordinates[i] < -s.BoundingRectangleSize {
			position.coordinates[i] = -s.BoundingRectangleSize
		} else if position.coordinates[i] > s.BoundingRectangleSize {
			position.coordinates[i] = s.BoundingRectangleSize
		}
	}
}

func NewSwarm(inertia, c1, c2 float64, GlobalBest Position, particles []*Particle, fitnessFunction FitnessFunction, boundarySize float64) *Swarm {
	return &Swarm{Inertia: inertia, ConstantOne: c1, ConstantTwo: c2, Particles: particles, GlobalBest: GlobalBest, FitnessFunc: fitnessFunction, BoundingRectangleSize: boundarySize}
}
