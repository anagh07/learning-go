package main

import "fmt"

type transformFunc func(int) int

func main() {
    numbers := []int{ 1, 3, 9, 21, 23, 27, 88 }
    doubledNumbers := transformNumbers(&numbers, double)
    tripledNumbers := transformNumbers(&numbers, triple)
    fmt.Println(doubledNumbers)
    fmt.Println(tripledNumbers)
}

func transformNumbers(numbers *[]int, double transformFunc) []int {
    newNumbers := []int{}
    
    for _, value := range *numbers {
        newNumbers = append(newNumbers, double(value))
    }

    return newNumbers
}

func double(num int) int {
    return 2 * num
}

func triple(num int) int {
    return 3 * num
}