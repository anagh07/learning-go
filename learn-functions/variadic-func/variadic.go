package main

import "fmt"

func main() {
    numbers := []int{ 1, 3, 5, 7, 11, 32 }
    sum := sumup(0, numbers...)
    anotherSum := sumup(100, 1, 4, 5, 6)
    fmt.Println(sum)
    fmt.Println(anotherSum)
}

func sumup(firstNum int, numbers ...int) int {
    sum := firstNum
    for _, val := range numbers {
        sum += val
    }
    return sum
}
