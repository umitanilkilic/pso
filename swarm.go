package pso

import (
	"math/rand"
)

type FitnessFunction func(Position) float64

type ConstraintFunction func(*Position)

type Swarm struct {
	// Inertia is the inertia of the swarm
	Inertia float64
	// ConstantOne is the social constant
	ConstantOne float64
	// ConstantTwo is the cognitive constant
	ConstantTwo float64
	// globalBest is the best position found by the swarm
	globalBest Position
	// Particles is a list of particles that the swarm will use to search the space
	Particles []Particle
	// FitnessFunc is the function that the swarm will use to evaluate the fitness of a position
	FitnessFunc FitnessFunction
	// Boundaries is the boundaries of the search space
	ConstraintFunc ConstraintFunction
}

func (s *Swarm) CalculateVelocity(particle *Particle) {
	inertiaPart := particle.GetVelocity().Multiply(s.Inertia)
	c1Part := particle.GetBestPosition().Subtract(*particle.GetPosition()).Multiply(s.ConstantOne * rand.Float64() * 2)
	c2Part := s.GetGlobalBest().Subtract(*particle.GetPosition()).Multiply(s.ConstantTwo * rand.Float64() * 2)
	newVelocity := inertiaPart.Add(c1Part).Add(c2Part)
	particle.UpdateVelocity(newVelocity)
}

func (s *Swarm) GetGlobalBest() Position {
	return s.globalBest
}

func (s *Swarm) UpdatePosition(particle *Particle) {
	newPosition := particle.GetPosition().Add(*particle.GetVelocity())
	if s.ConstraintFunc != nil {
		s.ConstraintFunc(&newPosition)
	}
	particle.UpdatePosition(newPosition)
}

func (s *Swarm) UpdateBestPositions(particle *Particle) {
	fitness := s.FitnessFunc(*particle.GetPosition())
	if fitness < s.FitnessFunc(*particle.GetBestPosition()) {
		particle.BestPosition = *particle.GetPosition()
	}
	if fitness < s.FitnessFunc(s.GetGlobalBest()) {
		s.globalBest = *particle.GetPosition()
	}
}

func NewSwarm(inertia, c1, c2 float64, particles []Particle, fitnessFunction FitnessFunction, constraintFunc ConstraintFunction) Swarm {
	swarm := Swarm{Inertia: inertia, ConstantOne: c1, ConstantTwo: c2, Particles: particles, FitnessFunc: fitnessFunction, ConstraintFunc: constraintFunc}
	if particles != nil || len(particles) != 0 {
		swarm.globalBest = *particles[0].GetPosition()
	}
	return swarm
}
