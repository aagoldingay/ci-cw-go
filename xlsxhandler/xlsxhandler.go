package xlsxhandler

import (
	"strconv"

	"github.com/tealeg/xlsx"
)

// WriteXLSXParams writes the values of a full execution to the parameter sheets of Data.xlsx
func WriteXLSXParams(revenue [][]float64, psoPopulation, aisPopulation, aisReplacement, aisClonesFactor int) {
	xl, err := xlsx.OpenFile("Data.xlsx")
	if err != nil {
		panic(err)
	}
	pso := xl.Sheets[0] // PSOParam
	ais := xl.Sheets[1] // AISParam

	psoRow := pso.AddRow()
	psoRow.AddCell().Value = strconv.Itoa(psoPopulation)
	psoRow.AddCell().Value = "0" // inertia add manually
	psoRow.AddCell().Value = "0" // cognitive add manually
	psoRow.AddCell().Value = "0" // social add manually
	for _, rev := range revenue[1] {
		cell := psoRow.AddCell()
		cell.Value = strconv.FormatFloat(rev, 'f', -1, 64) // consecutive prints of 3 tested seeds
	}

	aisRow := ais.AddRow()
	aisRow.AddCell().Value = strconv.Itoa(aisPopulation)
	aisRow.AddCell().Value = strconv.Itoa(aisClonesFactor)
	aisRow.AddCell().Value = strconv.Itoa(aisReplacement)
	aisRow.AddCell().Value = "0" // best fitness add manually
	for _, rev := range revenue[2] {
		cell := aisRow.AddCell()
		cell.Value = strconv.FormatFloat(rev, 'f', -1, 64) // consecutive prints of 3 tested seeds
	}

	err = xl.Save("Data.xlsx") // saves changes
	if err != nil {
		panic(err)
	}
}

// WriteXLSXRevenues writes values of all algorithm executions to Data.xlsx
func WriteXLSXRevenues(seeds []int64, randomRevenues, psoRevenues, aisRevenues [][]float64) {
	xl, err := xlsx.OpenFile("Data.xlsx")
	if err != nil {
		panic(err)
	}
	random := xl.Sheets[2]
	pso := xl.Sheets[3]
	ais := xl.Sheets[4]

	// record random
	for i := 0; i < len(seeds); i++ {
		randRow := random.AddRow()

		// first cell = seed
		randRow.AddCell().Value = strconv.FormatInt(seeds[i], 10)

		// next 100 for steps
		for j := 0; j < len(randomRevenues[i]); j++ {
			randRow.AddCell().Value = strconv.FormatFloat(randomRevenues[i][j], 'f', -1, 64)
		}
	}

	// record pso
	for i := 0; i < len(seeds); i++ {
		psoRow := pso.AddRow()

		// first cell = seed
		psoRow.AddCell().Value = strconv.FormatInt(seeds[i], 10)

		// next 100 for steps
		for j := 0; j < len(psoRevenues[i]); j++ {
			psoRow.AddCell().Value = strconv.FormatFloat(psoRevenues[i][j], 'f', -1, 64)
		}
	}

	// record ais
	for i := 0; i < len(seeds); i++ {
		aisRow := ais.AddRow()

		// first cell = seed
		aisRow.AddCell().Value = strconv.FormatInt(seeds[i], 10)

		// next 100 for steps
		for j := 0; j < len(aisRevenues[i]); j++ {
			aisRow.AddCell().Value = strconv.FormatFloat(aisRevenues[i][j], 'f', -1, 64)
		}
	}

	// save file and complete
	err = xl.Save("Data.xlsx")
	if err != nil {
		panic(err)
	}
}
