package main

import "github.com/aagoldingay/ci-cw-go/algorithms"

func main() {
	// algorithms.RandomSearch(20)
	// algorithms.PSOSearch(20, 10)
	algorithms.AISSearch(20, 20, 10, 8) //numGoods, numPopulation, replacement, cloneSizeFactor
}
