package algorithms

import (
	"fmt"
	"log"
	"math/rand"

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
// clones and mutates a population using elitism to generate better solutions
func AISSearch(numGoods, numPopulation, replacement, cloneSizeFactor int, trace bool, p *pp.PricingProblem) (float64, []float64) {
	revenueTrack := []float64{}
	population := ais.NewImmuneSystem(numGoods, numPopulation, replacement, cloneSizeFactor, p)
	fmt.Printf("best cell: %v\n", population.BestCell)

	for i := 0; i < 100; i++ {
		population.Update()
		fmt.Printf("[%v] best cell : %v\n", i+1, population.BestCell)

		// only track if required
		if trace {
			revenueTrack = append(revenueTrack, population.BestCell.Revenue)
		}
	}
	return population.BestCell.Revenue, revenueTrack
}

// PSOSearch is a CI algorithm approach to finding the highest possible revenue
// uses 'particles' to traverse the problem like a map, potentially encountering new, better results
func PSOSearch(numGoods, numParticles int, trace bool, p *pp.PricingProblem) (float64, []float64) {
	revenueTrack := []float64{}
	swarm := pso.NewSwarm(numGoods, numParticles, p)
	fmt.Printf("Particles created...\n")
	fmt.Printf("Best : %v | %v\n", swarm.BestPrices, swarm.BestRevenue)

	for i := 0; i < 100; i++ {
		swarm.Update()
		fmt.Printf("[%v] new best: prices : %v | revenue : %v\n", i+1, swarm.BestPrices, swarm.BestRevenue)

		// only track if required
		if trace {
			revenueTrack = append(revenueTrack, swarm.BestRevenue)
		}
	}
	return swarm.BestRevenue, revenueTrack
}

// RandomSearch is a heuristic method of attempting to find the highest possible revenue
// Approach : Create an array of random prices len(numGoods) and compare against the current best Revenue
// (This method was translated from the provided Java code)
func RandomSearch(numGoods int, trace bool, p *pp.PricingProblem) (float64, []float64) {
	revenueTrack := []float64{}
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

		// only track if required
		if trace {
			revenueTrack = append(revenueTrack, bestRevenue.revenue)
		}
	}
	fmt.Printf("Final best revenue : %v\n", bestRevenue)
	return bestRevenue.revenue, revenueTrack
}
