package main

import (
	"fmt"
	"time"

	"anagh.xyz/pricecalculator/fileutil"
	"anagh.xyz/pricecalculator/prices"
)

func main() {
	// Calculate program runtime
	startTime := time.Now().UnixMilli()

	taxrates := []float64{0.05, 0.1, 0.25}
	inputFilePath := "prices.txt"

	// Channels
	doneChannels := make([]chan bool, len(taxrates))
	errorChannels := make([]chan error, len(taxrates))

	for index, taxrate := range taxrates {
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxrate*100)
		fileManager := fileutil.New(inputFilePath, outputFilePath)
		job := prices.NewTaxInclPriceJob(fileManager, taxrate)
		// Create channels
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		go job.Process(doneChannels[index], errorChannels[index])
	}

	for index := range taxrates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err.Error())
			}
		case <-doneChannels[index]:
			fmt.Println("Done!")
		}
	}

	// Calculte program runtime
	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Program ran for %v milliseconds\n", duration)
}
