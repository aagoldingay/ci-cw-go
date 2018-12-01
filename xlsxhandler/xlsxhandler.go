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
