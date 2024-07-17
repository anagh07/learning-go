package prices

import (
	"fmt"

	"anagh.xyz/pricecalculator/conversion"
	"anagh.xyz/pricecalculator/fileutil"
)

type TaxInclPriceJob struct {
	TaxRate        float64
	InputPrices    []float64
	PricesAfterTax map[string]float64
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{4, 32, 25},
		TaxRate:     taxRate,
	}
}

func (job *TaxInclPriceJob) LoadData() {
	lines, err := fileutil.ReadLinesFromTextFile("prices.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	floatPrices, err := conversion.ConvertStringsToFloats(lines)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	job.InputPrices = floatPrices
}

func (job *TaxInclPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, inputPrice := range job.InputPrices {
		priceAfterTax := fmt.Sprintf("%.2f", (inputPrice * (1 + job.TaxRate)))
		result[fmt.Sprintf("%.2f", inputPrice)] = priceAfterTax
	}

	fmt.Println(result)
}
