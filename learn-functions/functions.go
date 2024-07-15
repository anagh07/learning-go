package main

import "fmt"

type transformFunc func(int) int

func main() {
    numbers := []int{ 1, 3, 9, 21, 23, 27, 88 }
    lessNumbers := []int{ 2, 4, 5 }
    // doubledNumbers := transformNumbers(&numbers, getTransformFunc(&numbers))
    // tripledNumbers := transformNumbers(&lessNumbers, getTransformFunc(&lessNumbers))
    doubledNumbers := transformNumbers(&numbers, func(number int) int {
        return 2 * number
    })
    tripledNumbers := transformNumbers(&lessNumbers, func (number int) int {
        return 3 * number
    })
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

func getTransformFunc(nums *[]int) transformFunc {
    if len(*nums) < 5 {
        return triple
    } else {
        return double
    }
}
