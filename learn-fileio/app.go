package main

import (
	"fmt"
	"strconv"

	"anagh.xyz/learninggo/utils"
	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	const fileName = "state.txt"
	var preferredDrink string
	name := "anagh"
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Preferred Drink: ")
	fmt.Scan(&preferredDrink)
	fmt.Println("My favorite drink is", preferredDrink)
	// fmt.Printf(`%s's favorite drink
	// is %s`, name, preferredDrink)

	// for i := 0; true; i++ {
	// 	var loopCheck bool
	// 	fmt.Scan(&loopCheck)
	// 	if !loopCheck {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	countString := utils.ReadFromFile(fileName)
	count, _ := strconv.ParseInt(countString, 10, 64)
	count++
	fmt.Println(count)
	utils.WriteToFile(fileName, fmt.Sprint(count))
	// Random phone
	fmt.Println(gofakeit.Phone())

	// Pointers and de-reference
	myCity := "Dhaka"
	var myCityPointer *string = &myCity
	fmt.Println(myCityPointer, *myCityPointer)
}
