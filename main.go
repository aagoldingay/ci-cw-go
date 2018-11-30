package main

import "github.com/aagoldingay/ci-cw-go/algorithms"

func main() {
	// algorithms.RandomSearch(20) //numGoods
	algorithms.PSOSearch(20, 25) //numGoods, numParticles
	// algorithms.AISSearch(20, 30, 10, 5) //numGoods, numPopulation, replacement, cloneSizeFactor
}
