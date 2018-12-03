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
// clones and mutates a population using elitism to generate better solutions
func AISSearch(numGoods, numPopulation, replacement, cloneSizeFactor int, trace bool, p *pp.PricingProblem) (float64, []float64) {
	revenueTrack := []float64{}
	population := ais.NewImmuneSystem(numGoods, numPopulation, replacement, cloneSizeFactor, p)
	fmt.Printf("Cells created...\n")
	//fmt.Printf("best cell: %v\n", population.BestCell)

	timeout := time.After(3 * time.Second)
	tick := time.Tick(5 * time.Millisecond)

	for { //for i := 0; i < 100; i++ {
		// if testing with time, uncomment switch statement, move inner case <-timeout out of for loop

		select {
		// timeout reached, stop running
		case <-timeout:
			fmt.Printf("Final best revenue : %v\n", population.BestCell)
			revenueTrack = append(revenueTrack, population.BestCell.Revenue) // adds 30th result
			return population.BestCell.Revenue, revenueTrack
		// tick reached, record data
		case <-tick:
			if trace {
				revenueTrack = append(revenueTrack, population.BestCell.Revenue)
			}
		// run procedure
		default:
			population.Update()
			//fmt.Printf("[%v] best cell : %v\n", i+1, population.BestCell) //uncomment on iterations
		}
	}
}

// PSOSearch is a CI algorithm approach to finding the highest possible revenue
// uses 'particles' to traverse the problem like a map, potentially encountering new, better results
func PSOSearch(numGoods, numParticles int, trace bool, p *pp.PricingProblem) (float64, []float64) {
	revenueTrack := []float64{}
	swarm := pso.NewSwarm(numGoods, numParticles, p)
	fmt.Printf("Particles created...\n")
	//fmt.Printf("Best : %v | %v\n", swarm.BestPrices, swarm.BestRevenue)

	timeout := time.After(3 * time.Second)
	tick := time.Tick(5 * time.Millisecond)

	for { //for i := 0; i < 100; i++ {
		// if testing with time, uncomment switch statement, move inner case <-timeout out of for loop

		select {
		// timeout reached, stop running
		case <-timeout:
			fmt.Printf("Final best revenue : %v\n", swarm.BestRevenue)
			revenueTrack = append(revenueTrack, swarm.BestRevenue) // adds 30th result
			return swarm.BestRevenue, revenueTrack
		// tick reached, record data
		case <-tick:
			if trace {
				revenueTrack = append(revenueTrack, swarm.BestRevenue)
			}
		// run procedure
		default:
			swarm.Update()
			//fmt.Printf("[%v] new best: prices : %v | revenue : %v\n", i+1, swarm.BestPrices, swarm.BestRevenue) // uncomment if steps		}
		}
	}
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

	timeout := time.After(3 * time.Second)
	tick := time.Tick(5 * time.Millisecond)

	for { //for i := 0; i < 100; i++ {
		// if testing with time, uncomment switch statement, move inner case <-timeout out of for loop

		select {
		// timeout reached, stop running
		case <-timeout:
			fmt.Printf("Final best revenue : %v\n", bestRevenue)
			revenueTrack = append(revenueTrack, bestRevenue.revenue) // adds 30th result
			return bestRevenue.revenue, revenueTrack
		// tick reached, record data
		case <-tick:
			if trace {
				revenueTrack = append(revenueTrack, bestRevenue.revenue)
			}
		// run procedure
		default:
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
				//fmt.Printf("New best revenue : %v \n", bestRevenue)
			}
		}
	}
}
