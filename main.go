package main

import (
	"fmt"

	"github.com/aagoldingay/ci-cw-go/algorithms"
	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
	"github.com/aagoldingay/ci-cw-go/xlsxhandler"
)

func main() {
	numGoods := 20
	//seeds := []int64{0, 38, 113} // simple, for parameter configuration
	seeds := []int64{0, 38, 113, 100, 50, 25, 75, 13, 55, 98, 187, 4, 12, 42, 66, 72, 30, 32, 10, 24, 49, 35, 88, 61, 19, 23, 14, 91, 102, 147}

	// runSingle(numGoods, seeds[0])
	runAll(numGoods, seeds)
}

func runSingle(numGoods int, seed int64) {
	p := pp.PricingProblem{}
	p = *p.MakeProblem(numGoods, seed, false) //courseworkInstance
	// p = *p.MakeProblem(numGoods, seed, true) //randomInstance

	rev, history := algorithms.RandomSearch(numGoods, true, &p) //numGoods
	fmt.Printf("rev : %v\nall : %v\n", rev, history)
	// algorithms.PSOSearch(numGoods, 25, false, &p) //numGoods, numParticles
	// algorithms.AISSearch(numGoods, 30, 10, 5, false, &p) //numGoods, numPopulation, replacement, cloneSizeFactor
}

func runAll(numGoods int, seeds []int64) {
	// configurable algorithm parameters
	psoPopulation := 20
	aisPopulation := 20
	aisReplacement := 10
	aisClonesFactor := 8

	// revenue trackers
	finalRevenues := [][]float64{}
	randomRevenues := [][]float64{}
	psoRevenues := [][]float64{}
	aisRevenues := [][]float64{}

	for i := 0; i < 3; i++ { // 3 algorithms
		finalRevenues = append(finalRevenues, make([]float64, len(seeds)))
	}

	for i := 0; i < len(seeds); i++ {
		p := pp.PricingProblem{}
		p = *p.MakeProblem(numGoods, seeds[i], false) //courseworkInstance
		// p = *p.MakeProblem(numGoods, true) //randomInstance

		// data structures for returned list of revenues per step of each process (for xlsx printing)
		var ran, pso, ais []float64

		fmt.Printf("----------\nRandom Search\n----------\n")
		finalRevenues[0][i], ran = algorithms.RandomSearch(numGoods, true, &p)
		randomRevenues = append(randomRevenues, ran)

		fmt.Printf("----------\nPSO\n----------\n")
		finalRevenues[1][i], pso = algorithms.PSOSearch(numGoods, psoPopulation, true, &p)
		psoRevenues = append(psoRevenues, pso)

		fmt.Printf("----------\nAIS\n----------\n")
		finalRevenues[2][i], ais = algorithms.AISSearch(numGoods, aisPopulation, aisReplacement, aisClonesFactor, true, &p) //numGoods, numPopulation, replacement, cloneSizeFactor
		aisRevenues = append(aisRevenues, ais)
	}
	fmt.Printf("%v\n", finalRevenues)

	// xlsx output
	// xlsxhandler.WriteXLSXParams(finalRevenues, psoPopulation, aisPopulation, aisReplacement, aisClonesFactor)
	xlsxhandler.WriteXLSXRevenues(seeds, randomRevenues, psoRevenues, aisRevenues)
}
