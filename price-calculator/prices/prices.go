package prices

import (
	"fmt"

	"anagh.xyz/pricecalculator/conversion"
	"anagh.xyz/pricecalculator/fileutil"
)

type TaxInclPriceJob struct {
	IOManager      fileutil.FileManager `json:"-"`
	TaxRate        float64              `json:"tax_rate"`
	InputPrices    []float64            `json:"prices"`
	PricesAfterTax map[string]string    `json:"price_after_tax"`
}

func NewTaxInclPriceJob(ioManager fileutil.FileManager, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		IOManager:   ioManager,
		InputPrices: []float64{4, 32, 25},
		TaxRate:     taxRate,
	}
}

func (job *TaxInclPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLinesFromTextFile()

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

	job.PricesAfterTax = result
	err := job.IOManager.WriteResult(job)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
