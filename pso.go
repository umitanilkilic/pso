package pso

type PSO struct {
	Swarm            *Swarm
	iteration        int
	currentIteration int
}

// runner is a function that executed after each iteration
type runner func(swarm *Swarm)

func NewPSO(swarm Swarm) PSO {
	return PSO{Swarm: &swarm}
}

func (pso *PSO) SetIterationCount(iteration int) {
	pso.iteration = iteration
}

func (pso *PSO) GetCurrentIteration() int {
	return pso.currentIteration
}

func (pso *PSO) GetSwarm() *Swarm {
	return pso.Swarm
}

func (pso *PSO) performNextIteration() {
	pso.Swarm.CalculateVelocity()
	pso.Swarm.UpdatePosition()
	pso.Swarm.UpdateBestPosition()
	pso.Swarm.UpdateGlobalBest()
}

func (pso *PSO) Optimize(executeFunc ...runner) {
	pso.Swarm.UpdateBestPosition()
	pso.Swarm.UpdateGlobalBest()
	for pso.currentIteration = 0; pso.currentIteration < pso.iteration; pso.currentIteration++ {
		pso.performNextIteration()
		if len(executeFunc) > 0 {
			executeFunc[0](pso.Swarm)
		}
	}
}
