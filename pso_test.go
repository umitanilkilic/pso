package pso

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {
	//easom function
	f := func(p Position) float64 {
		return -1 * (math.Cos(p.GetCoordinates()[0]) * math.Cos(p.GetCoordinates()[1]) * math.Exp(-1*(math.Pow(p.GetCoordinates()[0]-math.Pi, 2)+math.Pow(p.GetCoordinates()[1]-math.Pi, 2))))
	}
	//initialize particles
	particles := make([]*Particle, 0)
	for i := 0; i < 100; i++ {
		particles = append(particles, NewParticle(i, NewPosition(float64(rand.Float64()*10), float64(rand.Float64()*10)), NewPosition(float64(rand.Float64()*10), float64(rand.Float64()*10))))
	}
	//initialize swarm
	swarm := NewSwarm(0.5, 0.5, 0.5, particles, f, constraintFunc)
	pso := NewPSO(swarm)
	pso.Iteration = 1000
	//run pso
	pso.Optimize()

	fmt.Printf("Global best: %v\n", swarm.GetGlobalBest())
	fmt.Printf("Fitness: %v\n", f(*swarm.GetGlobalBest()))

}

func constraintFunc(position *Position) {
	for i := 0; i < position.dimensionSize; i++ {
		if position.coordinates[i] < -10 {
			position.coordinates[i] = -10
		} else if position.coordinates[i] > 10 {
			position.coordinates[i] = 10
		}
	}
}
