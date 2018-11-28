package ais

import (
	"math"
	"math/rand"
	"sort"

	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

// TCell models a price/revenue
type TCell struct {
	prices  []float64
	revenue float64
}

// ImmuneSystem is an object containing cells and parameter values
type ImmuneSystem struct {
	Cells                        []TCell
	BestCell                     TCell
	replacement, cloneSizeFactor int
	problem                      *pp.PricingProblem
	NormalisedRevenue            float64
}

const bestFitness = 8000.0

// NewImmuneSystem generates a new population of cells (prices and revenue)
func NewImmuneSystem(numGoods, numPopulation, replacement, cloneSizeFactor int, pr *pp.PricingProblem) *ImmuneSystem {
	is := new(ImmuneSystem)
	is.problem = pr
	is.replacement = replacement
	is.cloneSizeFactor = cloneSizeFactor

	population := make([]TCell, numPopulation)
	var totalFitness float64
	bestCell := TCell{}

	for i := 0; i < numPopulation; i++ {
		prices, rev := is.randomPrices(numGoods)
		population[i] = TCell{prices, rev}
		if i == 0 || bestCell.revenue < rev {
			bestCell = population[i]
		}
		totalFitness += population[i].revenue
	}
	is.Cells = population
	is.BestCell = bestCell
	is.NormalisedRevenue = bestCell.revenue / totalFitness
	return is
}

// Update acts as a step, and causes alterations on the population
func (is *ImmuneSystem) Update() {
	is.Cells = is.metaDynamics(is.clonalSelection())
	var totalFitness float64
	for i := 0; i < len(is.Cells); i++ {
		if i == 0 || is.Cells[i].revenue > is.BestCell.revenue {
			is.BestCell = is.Cells[i]
		}
		totalFitness += is.Cells[i].revenue
	}
	is.NormalisedRevenue = is.BestCell.revenue / totalFitness
}

func (is *ImmuneSystem) clonalSelection() []TCell {
	// create clones
	clones := [][]TCell{}
	numCopies := len(is.Cells) * is.cloneSizeFactor
	for i := 0; i < len(is.Cells); i++ {
		clonesOfIndex := make([]TCell, numCopies)
		for j := 0; j < numCopies; j++ {
			clonesOfIndex[j] = is.Cells[i]
			copy(clonesOfIndex[j].prices, is.Cells[i].prices) // deep copy array otherwise both will change
		}
		clones = append(clones, clonesOfIndex)
	}

	// mutation
	for i := 0; i < len(clones); i++ {
		for j := 0; j < len(clones[i]); j++ {
			mutationRate := math.Exp(-1 * clones[i][j].revenue / bestFitness)
			if rand.Float64() <= mutationRate {
				clones[i][j] = is.contiguousHyperMutation(clones[i][j].prices) // cant change
			}
		}
	}

	returnedClones := []TCell{}
	for i := 0; i < len(clones); i++ {
		returnedClones = append(returnedClones, clones[i]...)
	}
	return returnedClones
}

func (is *ImmuneSystem) contiguousHyperMutation(prices []float64) TCell {
	newPrices := []float64{}
	var hotspotA, hotspotB int

	for hotspotA == hotspotB {
		hotspotA = rand.Intn(len(prices) - 2)
		hotspotA = rand.Intn(len(prices) - 1)
	}
	if hotspotA > hotspotB {
		hotspotA, hotspotB = hotspotB, hotspotA
	}

	// add prices to newPrices, and when between hotspots in reverse order
	for i := 0; i < len(prices); i++ {
		if i < hotspotB && i > hotspotA {
			for i := hotspotB; i > hotspotA; i-- {
				newPrices = append(newPrices, prices[i])
				if i == hotspotA {
					i++
				}
			}
		} else {
			newPrices = append(newPrices, prices[i])
		}
	}
	for i := hotspotB; i > hotspotA; i-- {
		newPrices = append(newPrices, prices[i])
	}
	rev, _ := is.problem.Evaluate(newPrices)
	return TCell{newPrices, rev}
}

func (is *ImmuneSystem) metaDynamics(clones []TCell) []TCell {
	// combine population and clones
	newPopulation := []TCell{}
	newPopulation = append(newPopulation, is.Cells...)
	newPopulation = append(newPopulation, clones...)
	newPopulation = sortPopulation(newPopulation)

	newPopulation = newPopulation[:len(is.Cells)] //cuts population down to original population size

	//replace with random solutions
	for i := len(is.Cells) - is.replacement - 1; i < len(newPopulation); i++ {
		rp, rev := is.randomPrices(len(is.Cells[0].prices))
		newPopulation[i] = TCell{rp, rev}
	}
	return newPopulation
}

// randomPrices generates random prices that evaluated as valid by the PricingProblem
func (is *ImmuneSystem) randomPrices(numGoods int) ([]float64, float64) {
	prices := make([]float64, numGoods)
	for !is.problem.IsValid(prices) {
		for i := 0; i < numGoods; i++ {
			prices[i] = rand.Float64() * 10
		}
	}
	rev, _ := is.problem.Evaluate(prices)
	return prices, rev
}

// sortPopulation sorts the whole population of prices by the highest revenue first
func sortPopulation(p []TCell) (sortedPopulation []TCell) {
	sortedPopulation = make([]TCell, len(p))
	copy(sortedPopulation, p)
	sort.Slice(sortedPopulation, func(i, j int) bool {
		return sortedPopulation[i].revenue > sortedPopulation[j].revenue
	})
	return
}
