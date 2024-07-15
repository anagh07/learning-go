package main

import "fmt"

func main() {
    numbers := []int{ 1, 3, 9, 21, 23, 27, 88 }
    doubleFunc := createTransformFunc(2)
    tripleFunc := createTransformFunc(3)
    doubledNumbers := transformNumbers(&numbers, doubleFunc)
    tripledNumbers := transformNumbers(&numbers, tripleFunc)
    fmt.Println(doubledNumbers)
    fmt.Println(tripledNumbers)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
    newNumbers := []int{}
    
    for _, value := range *numbers {
        newNumbers = append(newNumbers, transform(value))
    }

    return newNumbers
}

func createTransformFunc(factor int) func(int) int {
    return func(num int) int {
        return factor * num
    }
}