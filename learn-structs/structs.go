package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, error := user.New(userFirstName, userLastName, userBirthdate)	
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	appUser.OutPutUserData()
	appUser.ClearUserNames()
	appUser.OutPutUserData()

	adminUser := user.NewAdmin("admin@test.com", "admin1234")
	adminUser.OutPutUserData()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
