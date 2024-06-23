# PSO (Particle Swarm Optimization)

## Installation

```bash
go get github.com/umitanilkilic/pso
```

## Usage

## Usage

To use the PSO library, you can import it into your Go project and start optimizing your problems. Here's an example of how to use it:

```go
package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/umitanilkilic/pso"
)

var spaceSize = 100.0
var psOptimizer pso.PSO

func main() {
	//easom function

	f := func(p pso.Position) float64 {
		return -1 * (math.Cos(p.GetCoordinates()[0]) * math.Cos(p.GetCoordinates()[1]) * math.Exp(-1*(math.Pow(p.GetCoordinates()[0]-math.Pi, 2)+math.Pow(p.GetCoordinates()[1]-math.Pi, 2))))
	}
	//initialize particles
	particles := createParticles()
	//initialize swarm
	swarm := pso.NewSwarm(0.5, 0.5, 0.5, particles, f, constraintFunc)
	psOptimizer = pso.NewPSO(&swarm)
	psOptimizer.SetIterationCount(10)
	//run pso
	psOptimizer.Optimize(printParticleInfo)

	fmt.Printf("Global best: %v\n", swarm.GetGlobalBest())
	fmt.Printf("Fitness: %v\n", f(swarm.GetGlobalBest()))

}

func constraintFunc(position *pso.Position) {
	for i := 0; i < position.GetDimension(); i++ {
		if position.GetCoordinates()[i] < -spaceSize {
			position.GetCoordinates()[i] = -spaceSize
		} else if position.GetCoordinates()[i] > spaceSize {
			position.GetCoordinates()[i] = spaceSize
		}
	}
}

func createParticles() []pso.Particle {
	particles := make([]pso.Particle, 0)
	for i := 0; i < 100; i++ {
		pos := pso.NewPosition(float64(rand.Float64()*spaceSize), float64(rand.Float64()*spaceSize))
		vel := pso.NewPosition(float64(rand.Float64()*spaceSize), float64(rand.Float64()*spaceSize))
		particles = append(particles, pso.NewParticle(i, pos, vel))
	}
	return particles
}

func printParticleInfo(particles *pso.Swarm) {
	iteration := psOptimizer.GetCurrentIteration()
	fmt.Printf("Iteration: %v, Global Best Position: %v\n", iteration, particles.GetGlobalBest())
}

```

Make sure to replace the objective function logic and the bounds with your own problem-specific implementation. You can also adjust the number of dimensions, particles, and iterations according to your needs.

For more information, please refer to the [PSO documentation](https://github.com/umitanilkilic/pso).



## Citation

If you use this code in your research, please cite the following paper:

```bibtex
@software{KILIC_Particle_Swarm_Optimization,
author = {KILIC, UMIT ANIL},
license = {MIT},
title = {{Particle Swarm Optimization Library}},
url = {https://github.com/umitanilkilic/pso}
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


