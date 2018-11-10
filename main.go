package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	numberOfGoods := 20
	rand.Seed(time.Now().UnixNano())
	p := PricingProblem{}
	//p = PricingProblem.MakeProblem(p, numberOfGoods, false) //courseworkInstance
	p = PricingProblem.MakeProblem(p, numberOfGoods, true) //randomInstance

	prices := make([]float64, numberOfGoods)
	newPrices := make([]float64, numberOfGoods)

	for i := 0; i < numberOfGoods; i++ {
		prices[i] = rand.Float64() * 10
	}

	bestRevenue, err := p.Evaluate(prices)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("Best revenue so far is %v\n", bestRevenue)

		for j := 0; j < numberOfGoods; j++ {
			newPrices[j] = rand.Float64() * 10
		}

		newRevenue, err := p.Evaluate(newPrices)
		if err != nil {
			log.Fatal(err)
		}
		if newRevenue > bestRevenue {
			copy(prices, newPrices)
			bestRevenue = newRevenue
		}
	}

	fmt.Printf("Final best revenue : %v\n", bestRevenue)
}
