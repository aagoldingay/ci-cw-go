package pso

import (
	"log"
	"math/rand"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

//movement weightings
const (
	inertia    = 0.721 // weighting of momentum maintained between steps
	cognitiveW = 1.12  // (default) 1.1193 // weighting towards personal best position
	socialW    = 1.1   // (default) 1.1193 // weighting towards global best position
)

// Particle is a struct representing a single particle entity
// Used to find optimal market pricing
type Particle struct {
	prices, velocity, bestPrices []float64
	currentRevenue, bestRevenue  float64
}

// Swarm models a population of Particles along with the current best prices and revenue
type Swarm struct {
	Particles   []*Particle
	BestPrices  []float64
	BestRevenue float64
	numGoods    int
	problem     *pp.PricingProblem
}

// NewSwarm generates a new population of Particles
func NewSwarm(numGoods int, numParticles int, pr *pp.PricingProblem) *Swarm {
	// define and populate new swarm
	sw := new(Swarm)
	sw.problem = pr
	sw.numGoods = numGoods
	sw.Particles = make([]*Particle, numParticles)

	// create the population of particles
	for i := 0; i < numParticles; i++ {
		sw.Particles[i] = sw.NewParticle(numGoods)
		if sw.Particles[i].currentRevenue > sw.BestRevenue {
			// assign values to best prices and revenue
			sw.BestPrices = sw.Particles[i].prices
			sw.BestRevenue = sw.Particles[i].currentRevenue
		}
	}
	return sw
}

// NewParticle generates and returns a new instanc of a particle
// With a randomly generated design
func (sw *Swarm) NewParticle(numGoods int) *Particle {
	//define and populate new particle
	p := new(Particle)
	p.prices = randomPrices(numGoods, sw.problem)
	p.velocity = initialVelocity(p.prices, randomPrices(numGoods, sw.problem))
	p.bestPrices = make([]float64, len(p.prices))
	copy(p.bestPrices, p.prices) //important to copy due to pass by reference
	p.currentRevenue = evaluatePrices(p.prices, sw.problem)
	p.bestRevenue = p.currentRevenue
	return p
}

// Update (Swarm) iterates over the population of particles to continue the progress of the swarm by one step
func (sw *Swarm) Update() {
	for i := 0; i < len(sw.Particles); i++ {
		sw.Particles[i].Update(sw.numGoods, sw.BestPrices, sw.problem)
		if sw.Particles[i].currentRevenue > sw.BestRevenue {
			// ensures the best result is updated as necessary
			sw.BestPrices = sw.Particles[i].prices
			sw.BestRevenue = sw.Particles[i].currentRevenue
		}
	}
}

// Update (Particle) handles the repositioning and evaluation of a particle
// param: gBestPrices passes information of the global best prices across a whole population of particles
func (p *Particle) Update(numGoods int, gBestPrices []float64, pr *pp.PricingProblem) {
	copy(p.velocity, calculateVelocity(p.velocity, p.prices, p.bestPrices, gBestPrices)) //important to copy due to pass by reference
	copy(p.prices, updatePosition(p.prices, p.velocity, pr))                             //important to copy due to pass by reference
	p.currentRevenue = evaluatePrices(p.prices, pr)
	if p.currentRevenue > p.bestRevenue {
		copy(p.bestPrices, p.prices) //important to copy due to pass by reference
		p.bestRevenue = p.currentRevenue
	}
}

// calculateVelocity calculates the movement properties ready for updating a Particle's position
// uses inertia, cognitiveW, socialW constants
func calculateVelocity(velocity, prices, pBestPrices, gBestPrices []float64) []float64 {
	newVelocity := make([]float64, len(velocity))
	for i := 0; i < len(velocity); i++ {
		r1, r2 := rand.Float64(), rand.Float64()
		newVelocity[i] = (inertia * velocity[i]) + (cognitiveW * r1 * (pBestPrices[i] - prices[i])) + (socialW * r2 * (gBestPrices[i] - prices[i]))
	}
	return newVelocity
}

// evaluatePrices calculates the revenue for the provided prices
func evaluatePrices(prices []float64, pr *pp.PricingProblem) float64 {
	revenue, err := pr.Evaluate(prices)
	if err != nil {
		log.Fatal(err)
	}
	return revenue
}

// initialVelocity calculates the initial velocity by comparing two randomly generated arrays of prices
func initialVelocity(firstPrices, secondPrices []float64) []float64 {
	velocity := make([]float64, len(firstPrices))

	for i := 0; i < len(firstPrices); i++ {
		velocity[i] = (secondPrices[i] - firstPrices[i]) / 2
	}
	return velocity
}

// randomPrices generates random prices that evaluated as valid by the PricingProblem
func randomPrices(numGoods int, pr *pp.PricingProblem) []float64 {
	prices := make([]float64, numGoods)
	for !pr.IsValid(prices) {
		for i := 0; i < numGoods; i++ {
			prices[i] = rand.Float64() * 10
		}
	}
	return prices
}

// updatePosition uses the velocity to update the location of the Particle
func updatePosition(prices, velocity []float64, pr *pp.PricingProblem) []float64 {
	newPrices := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		newPrices[i] = prices[i] + velocity[i]
	}
	return newPrices
}
