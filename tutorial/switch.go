package tutorial

import "fmt"

func main() {

	//city := "London"
	var city string
	fmt.Print("City: ")
	fmt.Scan(&city)

	switch city {
	case "New York":
		fmt.Println("1")
	case "London", "Mexico":
		fmt.Println("2")
	case "Berlin", "Moscow":
		fmt.Println(3)
	default:
		fmt.Println("Not valid city.")
	}

}
