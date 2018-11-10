package pricingproblem

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// PricingProblem contains information about prices and
type PricingProblem struct {
	priceResponseType           []int
	priceResponse, impact, bnds [][]float64
}

// MakeProblem instantiates a new PricingProblem
// n = number of Goods for the pricing problem
// random = whether to use a random seed, else seed = 0
func (p PricingProblem) MakeProblem(n int, random bool) PricingProblem {
	rand.Seed(time.Now().UnixNano())
	if !random {
		rand.Seed(0)
	}
	p.priceResponse = [][]float64{} // n by 2
	for i := 0; i < n; i++ {
		p.priceResponse = append(p.priceResponse, make([]float64, 2))
	}
	p.priceResponseType = make([]int, n)
	p.impact = [][]float64{} // n by n
	for i := 0; i < n; i++ {
		p.impact = append(p.impact, make([]float64, n))
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Setting up good {%v} with type: ", i)
		t := rand.Float64()
		if t <= 0.4 {
			// Linear
			p.priceResponseType[i] = 0
			p.priceResponse[i][0] = p.getRandomTotalDemand()
			p.priceResponse[i][1] = p.getRandomSatiatingPrice()
			fmt.Printf("L (%v/%v)\n", p.priceResponse[i][0], p.priceResponse[i][1])

		} else if t > 0.4 && t < 0.9 {
			// Constant Elasticity
			p.priceResponseType[i] = 1
			p.priceResponse[i][0] = p.getRandomTotalDemand()
			p.priceResponse[i][1] = p.getRandomElasticity()
			fmt.Printf("CE (%v/%v)\n", p.priceResponse[i][0], p.priceResponse[i][1])
		} else {
			// Fixed Demand
			p.priceResponseType[i] = 2
			p.priceResponse[i][0] = p.getRandomTotalDemand()
			fmt.Printf("FD (%v/%v)\n", p.priceResponse[i][0], p.priceResponse[i][1])
		}

		for j := 0; j < n; j++ {
			p.impact[i][j] = rand.Float64() * 0.1
		}
		p.impact[i][i] = 0.0
	}
	p.bnds = [][]float64{}
	for i := 0; i < len(p.priceResponse); i++ {
		p.bnds = append(p.bnds, make([]float64, 2))
	}
	dimBnd := []float64{0.01, 10.0}
	for i := 0; i < len(p.priceResponse); i++ {
		p.bnds[i] = dimBnd
	}
	return p
}

// Bounds returns the bnds variable of a PricingProblem struct
func (p PricingProblem) Bounds() [][]float64 {
	return p.bnds
}

// IsValid checks whether a vector of prices is valid
// A valid price vector is one in which all prices are at least 1p and at most £10.00
func (p PricingProblem) IsValid(prices []float64) bool {
	if len(prices) != len(p.Bounds()) {
		return false
	}
	// all antenna lie within problem bounds
	for i := 0; i < len(prices); i++ {
		if prices[i] < p.Bounds()[i][0] || prices[i] > p.Bounds()[i][1] {
			return false
		}
	}
	return true
}

// Evaluate gets the total revenue from pricing goods as given in parameter
func (p PricingProblem) Evaluate(prices []float64) (float64, error) {
	if len(prices) != len(p.Bounds()) {
		return 0.0, errors.New("PricingProblem::evaluate called on price array of the wrong size")
	}
	if !p.IsValid(prices) {
		return 0.0, nil
	}
	var revenue float64
	for i := 0; i < len(prices); i++ {
		revenue += float64(p.getDemand(i, prices)) * prices[i]
	}

	return math.Round(revenue*100.0) / 100.0, nil
}

// get the demand for good i at price p
func (p PricingProblem) getDemand(i int, prices []float64) int {
	demand := p.getGoodDemand(i, prices[i]) + p.getResidualDemand(i, prices)

	// Second sanity check - still cannot have more demand than the market holds
	if float64(demand) > p.priceResponse[i][0] {
		demand = int(math.Round(p.priceResponse[i][0]))
	}
	return demand
}

func (p PricingProblem) getGoodDemand(i int, price float64) int {
	var demand float64
	switch p.priceResponseType[i] {
	case 0: // Linear
		demand = p.priceResponse[i][0] - ((p.priceResponse[i][0] / p.priceResponse[i][1]) * price)
		break
	case 1: // Constant Elasticity
		demand = p.priceResponse[i][0] / (math.Pow(price, p.priceResponse[i][1]))
		break
	case 2: // Fixed Demand
		demand = p.priceResponse[i][0]
		break
	default:
		fmt.Println("Error! Incorrect price response curve specified")
	}

	// Sanity checks - cannot have more demand than market holds
	if demand > p.priceResponse[i][0] {
		demand = math.Round(p.priceResponse[i][0])
	}
	// or less than 0 demand
	if demand < 0 {
		demand = 0
	}
	return int(math.Round(demand))
}

func (p PricingProblem) getResidualDemand(i int, prices []float64) int {
	var demand float64
	for j := 0; j < len(p.priceResponse); j++ {
		if i != j {
			demand += float64(p.getGoodDemand(j, prices[j])) * p.impact[j][i]
		}
	}
	return int(math.Round(demand))
}

func (p PricingProblem) getRandomTotalDemand() float64 {
	return rand.Float64() * 100
}

func (p PricingProblem) getRandomSatiatingPrice() float64 {
	return rand.Float64() * 10
}

func (p PricingProblem) getRandomElasticity() float64 {
	return rand.Float64()
}
