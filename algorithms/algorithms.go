package algorithms

import (
	"fmt"
	"log"
	"math/rand"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

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

	bestRevenue, err := p.Evaluate(prices)
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
		if newRevenue > bestRevenue {
			copy(prices, newPrices)
			bestRevenue = newRevenue
			fmt.Printf("New best revenue : %v | %v\n", newRevenue, prices)
		}
	}
	fmt.Printf("Final best revenue : %v | %v\n", bestRevenue, prices)
}
