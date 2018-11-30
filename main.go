package main

import "github.com/aagoldingay/ci-cw-go/algorithms"

func main() {
	// algorithms.RandomSearch(20) //numGoods
	algorithms.PSOSearch(20, 10) //numGoods, numParticles
	// algorithms.AISSearch(20, 20, 10, 10) //numGoods, numPopulation, replacement, cloneSizeFactor
}
