package prices

import "fmt"

type TaxIncluldPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncluldPrices map[string]float64
}

func (job TaxIncluldPrice) Process() {
	result := make(map[string]float64)
	for _, p := range job.InputPrices {
		result[fmt.Sprintf("%.2f", p)] = p + (1 * job.TaxRate)
	}

}

func NewTaxIncluldPrice(taxRate float64, inputPrices []float64) *TaxIncluldPrice {
	return &TaxIncluldPrice{
		TaxRate:          taxRate,
		InputPrices:      []float64{20.0, 30.0, 40.0},
		TaxIncluldPrices: make(map[string]float64),
	}
}
