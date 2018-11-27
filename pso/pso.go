package pso

import (
	"log"
	"math/rand"
	"time"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

//movement weightings
const (
	inertia    = 0.721  // weighting of momentum maintained between steps
	cognitiveW = 1.1193 // weighting towards personal best position
	socialW    = 1.1193 // weighting towards global best position
)

// Particle is a struct representing a single particle entity
// Will attempt to find optimal market pricing
type Particle struct {
	Prices, velocity, BestPrices []float64
	CurrentRevenue, BestRevenue  float64
}

// NewParticle generates and returns a new instanc of a particle
// With a randomly generated design
func NewParticle(numGoods int, pr *pp.PricingProblem) *Particle {
	p := new(Particle)
	p.Prices = randomPrices(numGoods, pr)
	p.velocity = initialVelocity(p.Prices, randomPrices(numGoods, pr))
	p.BestPrices = make([]float64, len(p.Prices))
	copy(p.BestPrices, p.Prices) //important to copy due to pass by reference
	p.CurrentRevenue = evaluatePrices(p.Prices, pr)
	p.BestRevenue = p.CurrentRevenue
	return p
}

// Update handles the repositioning and evaluation of a particle
// param: gBestPrices passes information of the global best prices across a whole population of particles
func (p *Particle) Update(numGoods int, gBestPrices []float64, pr *pp.PricingProblem) {
	copy(p.velocity, calculateVelocity(p.velocity, p.Prices, p.BestPrices, gBestPrices))
	copy(p.Prices, updatePosition(p.Prices, p.velocity, numGoods, pr))
	p.CurrentRevenue = evaluatePrices(p.Prices, pr)
	if p.CurrentRevenue > p.BestRevenue {
		copy(p.BestPrices, p.Prices)
		p.BestRevenue = p.CurrentRevenue
	}
}

// calculateVelocity calculates the movement properties ready for updating a Particle's position
// uses inertia, cognitiveW, socialW constants
func calculateVelocity(velocity, prices, pBestPrices, gBestPrices []float64) []float64 {
	newVelocity := make([]float64, len(velocity))
	for i := 0; i < len(velocity); i++ {
		rand.Seed(time.Now().UnixNano())
		r1, r2 := rand.Float64(), rand.Float64()
		newVelocity[i] = (inertia * velocity[i]) + (cognitiveW * r1 * (pBestPrices[i] - prices[i])) + (socialW * r2 * (gBestPrices[i] - prices[i]))
	}
	return newVelocity
}

func evaluatePrices(prices []float64, pr *pp.PricingProblem) float64 {
	revenue, err := pr.Evaluate(prices)
	if err != nil {
		log.Fatal(err)
	}
	return revenue
}

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

func updatePosition(prices, velocity []float64, numGoods int, pr *pp.PricingProblem) []float64 {
	newPrices := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		newPrices[i] = prices[i] + velocity[i]
	}
	return newPrices
}
