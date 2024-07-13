package main

import "fmt"

func main() {
	prices := [5]float32{10.2, 21, 1.234, 10, 27}
	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(prices)
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)

	var names [3]string = [3]string{"Anagh"}
	names[1] = "Mehran"
	names[2] = "N/A"
	fmt.Println(names)
	twoNames := names[0:2]
	fmt.Println(twoNames)

	// slices are references to the original array
	// if slice is edited, array is edited
	firstTwoPrices := prices[:2]
	lastTwoPrices := prices[3:]
	firstFourPrices := firstTwoPrices[:4]
	fmt.Printf("firstTwoPrices: %v\n", firstTwoPrices)
	fmt.Printf("lastTwoPrices: %v\n", lastTwoPrices)
	fmt.Printf("firstFourPrices: %v\n", firstFourPrices)
	firstTwoPrices[1] = 77
	fmt.Printf("After edit prices: %v\n", prices)
	newLastTwoPrices := lastTwoPrices[2:] // empty
	fmt.Println(newLastTwoPrices)

	// Dynamic arrays
	priceHistory := []float64{9.123, 71, 32.23}
	priceHistory = append(priceHistory, 5.51)
	fmt.Printf("Dynamic list: %v\n", priceHistory)

	// merge slices
	newPrices := []float32{.23, 5.12, 9.91}
	newPrices = append(newPrices, prices[0:]...) // spread operator unpacks the array
	fmt.Printf("Merged lists: %v\n", newPrices)

	// Pre-allot list size with make
	restaurants := make([]string, 2, 6)
	restaurants[0] = "3Amigos"
	fmt.Printf("Size: %v\tCapacity: %v\n", len(restaurants), cap(restaurants))
	restaurants[1] = "Miran"
	// restaurants[2] = "Nandos"
	restaurants = append(restaurants, "Nandos")
	fmt.Println(restaurants)

	// Loop through list
	for index, value := range restaurants {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}
}
