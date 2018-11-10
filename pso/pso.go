package pso

import (
	"log"
	"math/rand"

	"github.com/aagoldingay/ci-cw-go/pricingproblem"
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
	prices, velocity, bestPrices []float64
	currentRevenue, bestRevenue  float64
}

// NewParticle generates and returns a new instanc of a particle
// With a randomly generated design
func NewParticle(numberOfGoods int, pp *pricingproblem.PricingProblem) *Particle {
	p := new(Particle)
	p.prices = randomPrices(numberOfGoods)
	p.velocity = initialVelocity(p.prices, randomPrices(numberOfGoods))
	p.bestPrices = make([]float64, len(p.prices))
	copy(p.bestPrices, p.prices) //important to copy due to pass by reference
	p.currentRevenue = evaluatePrices(numberOfGoods, p.prices, pp)
	p.bestRevenue = p.currentRevenue
	return p
}

func evaluatePrices(numberOfGoods int, prices []float64, pp *pricingproblem.PricingProblem) float64 {
	revenue, err := pp.Evaluate(prices)
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

func randomPrices(numberOfGoods int) []float64 {
	prices := make([]float64, numberOfGoods)
	for i := 0; i < numberOfGoods; i++ {
		prices[i] = rand.Float64() * 10
	}
	return prices
}
