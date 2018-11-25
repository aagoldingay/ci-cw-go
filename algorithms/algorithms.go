package algorithms

import (
	"fmt"
	"log"
	"math/rand"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

// Revenue is a struct acting as a payload to access prices and revenues
// intended to store best revenue
type Revenue struct {
	revenue float64
	prices  []float64
}

// PSOSearch is a CI algorithm approach to finding the highest possible revenue
func PSOSearch(numGoods, numParticles int) {
	p := pp.PricingProblem{}
	//p = *p.MakeProblem(numGoods, false) //courseworkInstance
	p = *p.MakeProblem(numGoods, true) //randomInstance
}

// RandomSearch is a heuristic method of attempting to find the highest possible revenue
// Approach : Create an array of random prices len(numGoods) and compare against the current best Revenue
func RandomSearch(numGoods int) {
	p := pp.PricingProblem{}
	//p = *p.MakeProblem(numGoods, false) //courseworkInstance
	p = *p.MakeProblem(numGoods, true) //randomInstance

	prices := make([]float64, numGoods)
	newPrices := make([]float64, numGoods)

	for i := 0; i < numGoods; i++ {
		prices[i] = rand.Float64() * 10
	}

	bRevenue, err := p.Evaluate(prices)
	bestRevenue := Revenue{bRevenue, prices}
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
