package tutorial

import "fmt"

func main() {

	var conferenceName = "Go conf"

	const conferenceTickets int = 50

	var remainingTickets uint = 30

	fmt.Printf("Welcome here: %v champs!\n", conferenceName)
	fmt.Printf("We have total: %d\n", conferenceTickets)
	fmt.Println("There still available", remainingTickets)
	fmt.Println("You can buy tickets here")

	// var userName string
	// var userAge int
	// // comments
	// userName = "Ben"
	// userAge = 30

	// fmt.Printf("Your name %v and your age %v years!\n", userName, userAge)

	fmt.Println("===================================")

	fmt.Printf("conferenceName type: %T\n", conferenceName)

	fmt.Println("===================================")

	fmt.Println("syntax sugar:")
	randomVar := "Python better than Go"
	//var randomVar string = "Python better than Go"
	fmt.Println(randomVar)

	fmt.Println("===================================")
	fmt.Println("Pointer")

	var userName string
	var userAge int
	var ticketsN uint

	// fmt.Println(&userName)
	// fmt.Println(userAge)
	fmt.Print("What's your name ")
	fmt.Scan(&userName)
	fmt.Print("how old are you ")
	fmt.Scan(&userAge)
	fmt.Print("How mane tickets ")
	fmt.Scan(&ticketsN)
	// pointer &userName - it's a memory location
	fmt.Printf("Your name %v and your age %v years!\n", userName, userAge)
	fmt.Println("===================================")

	remainingTickets = remainingTickets - ticketsN
	fmt.Print("Ramaining tickets", remainingTickets)

	fmt.Println("===================================")
	fmt.Println("Arrays, Slices")
	// arrays have fixed size and type
	//var bookings = [50]string{"hello", "serge", "mymail@mail.com"}
	//var bookings = [50]string{}

	var bookings [50]string
	bookings[0] = "hello"
	bookings[10] = "serge"

}
