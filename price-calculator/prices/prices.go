package prices

import "fmt"

type TaxInclPriceJob struct {
    TaxRate float64
    InputPrices []float64
    PricesAfterTax map[string]float64
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
    return &TaxInclPriceJob{
        InputPrices: []float64{4, 32, 25},
        TaxRate: taxRate,
    }
}

func (job TaxInclPriceJob) Process() {
    result := make(map[string]float64)
    for _, inputPrice := range job.InputPrices {
        result[fmt.Sprintf("%.2f", inputPrice)] = inputPrice * (1 + job.TaxRate)
    }

    fmt.Println(result)
}