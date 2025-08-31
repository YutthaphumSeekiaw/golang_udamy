package main

import (
	"fmt"
	"log"
	"pricecal/filemanager"
)

func main() {
	lines, err := filemanager.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

	// taxRates := []float64{0.05, 0.1, 0.15, 0.2}

	// result := make(map[float64][]float64)

	// for _, taxRate := range taxRates {
	// 	pricesJob := prices.NewTaxIncluldPrice(taxRate)
	// }

	// fmt.Println(result)
}
