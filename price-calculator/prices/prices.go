package prices

import (
	"fmt"

	"anagh.xyz/pricecalculator/conversion"
	"anagh.xyz/pricecalculator/iomanager"
)

type TaxInclPriceJob struct {
	IOManager      iomanager.IOManager `json:"-"`
	TaxRate        float64             `json:"tax_rate"`
	InputPrices    []float64           `json:"prices"`
	PricesAfterTax map[string]string   `json:"price_after_tax"`
}

func NewTaxInclPriceJob(io iomanager.IOManager, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		IOManager:   io,
		InputPrices: []float64{4, 32, 25},
		TaxRate:     taxRate,
	}
}

func (job *TaxInclPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	floatPrices, err := conversion.ConvertStringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = floatPrices
	return nil
}

func (job *TaxInclPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, inputPrice := range job.InputPrices {
		priceAfterTax := fmt.Sprintf("%.2f", (inputPrice * (1 + job.TaxRate)))
		result[fmt.Sprintf("%.2f", inputPrice)] = priceAfterTax
	}

	job.PricesAfterTax = result
	err = job.IOManager.WriteResult(job)
	if err != nil {
		return err
	}
	return nil
}
