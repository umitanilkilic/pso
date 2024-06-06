package pso

type PSO struct {
	swarm            *Swarm
	iteration        int
	currentIteration int
}

// runner is a function that executed after each iteration
type runner func(swarm *Swarm)

func NewPSO(sw Swarm) PSO {
	return PSO{swarm: &sw}
}

func (pso *PSO) SetIterationCount(iteration int) {
	pso.iteration = iteration
}

func (pso *PSO) GetCurrentIteration() int {
	return pso.currentIteration
}

func (pso *PSO) GetSwarm() *Swarm {
	return pso.swarm
}

func (pso *PSO) performNextIteration() {
	for i := 0; i < len(pso.swarm.Particles); i++ {
		particle := &pso.swarm.Particles[i]
		pso.swarm.CalculateVelocity(particle)
		pso.swarm.UpdatePosition(particle)
		pso.swarm.UpdateBestPositions(particle)
	}
}

func (pso *PSO) Optimize(executeFunc ...runner) {
	for i := 0; i < len(pso.swarm.Particles); i++ {
		particle := &pso.swarm.Particles[i]
		pso.swarm.UpdateBestPositions(particle)
	}
	for pso.currentIteration = 0; pso.currentIteration < pso.iteration; pso.currentIteration++ {
		pso.performNextIteration()
		if len(executeFunc) > 0 {
			executeFunc[0](pso.swarm)
		}
	}
}
