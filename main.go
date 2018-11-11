package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

func main() {
	numberOfGoods := 20
	rand.Seed(time.Now().UnixNano())
	p := pp.PricingProblem{}
	//p = p.MakeProblem(numberOfGoods, false) //courseworkInstance
	p = p.MakeProblem(numberOfGoods, true) //randomInstance

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
			fmt.Printf("New best revenue : %v\n", newRevenue)
		}
	}
	fmt.Printf("Final best revenue : %v\n", bestRevenue)
}
