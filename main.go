package main

import (
	"begining/helper"
	"begining/other_user"
	"sync"
)

//var generalName string = "Real name" //package level
var wg = sync.WaitGroup{}

func main() {
	userName := helper.GetName()

	switch userName {
	case "Serge":
		greetUser("Main programm", userName)
	default:
		other_user.Greet_non_User("Main programm", userName)
	}

	wg.Add(10) // wait group
	for i := 1; i <= 10; i++ {
		go helper.SendMessage(&wg)
	}
	wg.Wait()
	// time.Sleep(10 * time.Second) // dirty solution (workaround)
}
