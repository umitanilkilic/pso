# PSO (Particle Swarm Optimization)

## Installation

```bash
go get github.com/umitanilkilic/pso
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/umitanilkilic/pso"
)

func main() {
	//easom function
	f := func(p Position) float64 {
		return -1 * (math.Cos(p.GetCoordinates()[0]) * math.Cos(p.GetCoordinates()[1]) * math.Exp(-1*(math.Pow(p.GetCoordinates()[0]-math.Pi, 2)+math.Pow(p.GetCoordinates()[1]-math.Pi, 2))))
	}
	//initialize particles
	particles := make([]*Particle, 0)
	for i := 0; i < 10; i++ {
		particles = append(particles, NewParticle(i, NewPosition(float64(rand.Float64()*10), float64(rand.Float64()*10)), NewPosition(float64(rand.Float64()*10), float64(rand.Float64()*10))))
	}
	//initialize swarm
	swarm := CreateSwarm(0.5, 0.5, 0.5, NewPosition(0, 0), particles, f, 10)
	pso := NewPSO(swarm)
	pso.Iteration = 1000
	//run pso
	pso.Optimize(printParticleInfo)

	fmt.Printf("Global best: %v\n", swarm.GetGlobalBest())
	fmt.Printf("Fitness: %v\n", f(*swarm.GetGlobalBest()))

}

func printParticleInfo(particles []*Particle) {
	for _, v := range particles {
		fmt.Printf("Particle ID: %v\n", v.ID)
		fmt.Printf("Position: %v\n", *v.GetPosition())
	}
}

```

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


