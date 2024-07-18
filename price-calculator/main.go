package main

import (
	"fmt"

	"anagh.xyz/pricecalculator/fileutil"
	"anagh.xyz/pricecalculator/prices"
)

func main() {
	taxrates := []float64{0.05, 0.1, 0.25}
	inputFilePath := "prices.txt"

	for _, taxrate := range taxrates {
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxrate*100)
		fileManager := fileutil.New(inputFilePath, outputFilePath)
		job := prices.NewTaxInclPriceJob(fileManager, taxrate)
		// cmd := cmdmanager.New()
		// job := prices.NewTaxInclPriceJob(cmd, taxrate)
		err := job.Process()
		if err != nil {
			fmt.Println("error completing job")
			fmt.Println(err.Error())
			return
		}
	}
}
