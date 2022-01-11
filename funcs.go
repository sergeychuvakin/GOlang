package main

import (
	"begining/helper"
	"fmt"
)

func greetUser(confName string, userName string) {

	// userName := getName()
	userAge := helper.CaclYears()

	fmt.Printf("Welcome to special tutorial which we call:\n%s\n", confName)

	fmt.Printf("Hello %v!\nNice to meet you!\nYour age: %d\n", userName, userAge)

}
