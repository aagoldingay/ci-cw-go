package algorithms

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aagoldingay/ci-cw-go/ais"
	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
	"github.com/aagoldingay/ci-cw-go/pso"
)

// Revenue is a struct acting as a payload to access prices and revenues
// intended to store best revenue
type Revenue struct {
	prices  []float64
	revenue float64
}

// AISSearch is a CI algorithm approach to finding the highest possible revenue
func AISSearch(numGoods, numPopulation, replacement, cloneSizeFactor int) {
	p := pp.PricingProblem{}
	p = *p.MakeProblem(numGoods, false) //courseworkInstance
	// p = *p.MakeProblem(numGoods, true) //randomInstance
	population := ais.NewImmuneSystem(numGoods, numPopulation, replacement, cloneSizeFactor, &p)
	fmt.Printf("best cell: %v\n", population.BestCell)

	for i := 0; i < 100; i++ {
		population.Update()
		fmt.Printf("[%v] best cell : %v\n", i+1, population.BestCell)
	}
}

// PSOSearch is a CI algorithm approach to finding the highest possible revenue
func PSOSearch(numGoods, numParticles int) {
	p := pp.PricingProblem{}
	p = *p.MakeProblem(numGoods, false) //courseworkInstance
	// p = *p.MakeProblem(numGoods, true) //randomInstance
	bestRev := Revenue{[]float64{}, 0.0}
	particles := []*pso.Particle{}
	fmt.Printf("Creating particles ...\n")
	for i := 0; i < numParticles; i++ {
		particles = append(particles, pso.NewParticle(numGoods, &p))
		if particles[i].BestRevenue > bestRev.revenue {
			bestRev.prices = particles[i].BestPrices
			bestRev.revenue = particles[i].BestRevenue
		}
	}
	fmt.Printf("Particles created, commencing search...\n")

	timeout := time.Now().Add(10 * time.Second)

	for {
		if time.Now().After(timeout) {
			fmt.Println("Time limit reached. Ended search")
			break
		}

		for i := 0; i < len(particles); i++ {
			particles[i].Update(numGoods, bestRev.prices, &p)

			if particles[i].BestRevenue > bestRev.revenue {
				bestRev.prices = particles[i].BestPrices
				bestRev.revenue = particles[i].BestRevenue
				fmt.Printf("New best: prices : %v | best revenue : %v\n", bestRev.prices, bestRev.revenue)
			}
		}
	}
}

// RandomSearch is a heuristic method of attempting to find the highest possible revenue
// Approach : Create an array of random prices len(numGoods) and compare against the current best Revenue
func RandomSearch(numGoods int) {
	p := pp.PricingProblem{}
	p = *p.MakeProblem(numGoods, false) //courseworkInstance
	// p = *p.MakeProblem(numGoods, true) //randomInstance

	prices := make([]float64, numGoods)
	newPrices := make([]float64, numGoods)

	for i := 0; i < numGoods; i++ {
		prices[i] = rand.Float64() * 10
	}

	bRevenue, err := p.Evaluate(prices)
	bestRevenue := Revenue{prices, bRevenue}
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {

		for j := 0; j < numGoods; j++ {
			newPrices[j] = rand.Float64() * 10
		}

		newRevenue, err := p.Evaluate(newPrices)
		if err != nil {
			log.Fatal(err)
		}
		if newRevenue > bestRevenue.revenue {
			copy(bestRevenue.prices, newPrices)
			bestRevenue.revenue = newRevenue
			fmt.Printf("New best revenue : %v \n", bestRevenue)
		}
	}
	fmt.Printf("Final best revenue : %v\n", bestRevenue)
}
