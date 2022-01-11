package tutorial

import "fmt"

func main() {

	var bookings = [50]string{"1", "2", "3"}
	//var bookings [50]string
	bookings[0] = "hello"
	bookings[10] = "serge"
	fmt.Println(bookings)
	fmt.Printf("type of booking: %T\n", bookings)
	fmt.Printf("Length of booking: %v\n", len(bookings))

	fmt.Println("===== Slices =====")
	fmt.Println("slices have dynamic length")

	//var bookingS []string
	//bookingS := []string{}
	var bookingS = []string{}
	bookingS = append(bookingS, "5", "6")
	fmt.Println(bookingS)
	fmt.Printf("Length of bookingS: %v\n", len(bookingS))
}
