package main

import (
	"fmt"

	"github.com/aagoldingay/ci-cw-go/algorithms"
	pp "github.com/aagoldingay/ci-cw-go/pricingproblem"
)

func main() {
	numGoods := 20
	seeds := []int64{0, 38, 113}

	//runSingle(numGoods, seeds[0])
	runAll(numGoods, seeds)
}

func runSingle(numGoods int, seed int64) {
	p := pp.PricingProblem{}
	p = *p.MakeProblem(numGoods, seed, false) //courseworkInstance
	// p = *p.MakeProblem(numGoods, seed, true) //randomInstance

	// algorithms.RandomSearch(numGoods, &p) //numGoods
	algorithms.PSOSearch(numGoods, 25, &p) //numGoods, numParticles
	// algorithms.AISSearch(numGoods, 30, 10, 5, &p) //numGoods, numPopulation, replacement, cloneSizeFactor
}

func runAll(numGoods int, seeds []int64) {
	psoPopulation := 20

	aisPopulation := 20
	aisReplacement := 10
	aisClonesFactor := 8

	revenues := [][]float64{}

	for i := 0; i < 3; i++ {
		revenues = append(revenues, make([]float64, len(seeds)))
	}

	for i := 0; i < len(seeds); i++ {
		p := pp.PricingProblem{}
		p = *p.MakeProblem(numGoods, seeds[i], false) //courseworkInstance
		// p = *p.MakeProblem(numGoods, true) //randomInstance

		fmt.Printf("----------\nRandom Search\n----------\n")
		revenues[0][i] = algorithms.RandomSearch(numGoods, &p)

		fmt.Printf("----------\nPSO\n----------\n")
		revenues[1][i] = algorithms.PSOSearch(numGoods, psoPopulation, &p)
		fmt.Printf("----------\nAIS\n----------\n")
		revenues[2][i] = algorithms.AISSearch(numGoods, aisPopulation, aisReplacement, aisClonesFactor, &p) //numGoods, numPopulation, replacement, cloneSizeFactor
	}
	fmt.Printf("%v\n", revenues)

	//xlsxhandler.WriteXLSXParams(revenues, psoPopulation, aisPopulation, aisReplacement, aisClonesFactor)
}
