package pso

type PSO struct {
	Swarm     *Swarm
	Iteration int
}

// runner is a function that executed after each iteration
type runner func(swarm *Swarm)

func NewPSO(swarm *Swarm) *PSO {
	return &PSO{Swarm: swarm}
}

func (pso *PSO) performNextIteration() {
	pso.Swarm.CalculateVelocity()
	pso.Swarm.UpdatePosition()
	pso.Swarm.UpdateBestPosition()
	pso.Swarm.UpdateGlobalBest()
}

func (pso *PSO) Optimize(executeFunc ...runner) (bestPosition *Position) {
	pso.Swarm.UpdateBestPosition()
	pso.Swarm.UpdateGlobalBest()
	for i := 0; i < pso.Iteration; i++ {
		pso.performNextIteration()
		if len(executeFunc) > 0 {
			executeFunc[0](pso.Swarm)
		}
	}
	bestPosition = pso.Swarm.GetGlobalBest()
	return
}
