package other_user

import (
	"begining/helper"
	"fmt"
	"log"
)

// in packages all the functions should starts from capital letter
func Greet_non_User(confName string, userName string) {

	// userName := getName()
	userAge := helper.CaclYears()

	fmt.Printf("Welcome to special tutorial which we call:\n%s\n", confName)

	fmt.Printf("Hello %v!\nNice to meet you!\nYour age: %d\n", userName, userAge)

	reject()
}

func reject() {
	log.Print("Actually, you are not supposed to be here")
}
