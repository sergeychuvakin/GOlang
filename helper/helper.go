package helper

import (
	"fmt"
	"sync"
	"time"
)

func CaclYears() uint {

	var birth uint

	fmt.Print("What's your birth year? ")
	fmt.Scan(&birth)

	return 2022 - birth
}

func GetName() string {
	var userName string

	fmt.Print("What's your name? ")
	fmt.Scan(&userName)
	return userName
}

func SendMessage(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("#############################")
	fmt.Println("Hey Heys Hey")
	fmt.Println("#############################")
	wg.Done()
}
