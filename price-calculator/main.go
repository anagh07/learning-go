package main

import "anagh.xyz/pricecalculator/prices"

func main() {
    taxrates := []float64{0.05, 0.1, 0.25}
    
    for _, value := range taxrates {
        job := prices.NewTaxInclPriceJob(value)
        job.Process()
    }
}
