package main

import (
	"fmt"
)

func print(values ...interface{}) {

	for _, i := range values {
		fmt.Print(i, " ")

	}
	fmt.Print("\n")

}

func main() {

	s := "hello" // type string
	//s1 := "world"
	// t := "bye"        // type string
	// u := 44           // type int
	// v := [2]int{1, 2} // type array

	var pointer_s *string // type pointer
	pointer_s = &s
	//pointer_s = "ciao"

	print(pointer_s)
	print(*pointer_s == s)

	// print("s: ", s)
	//print(s1)

	// fmt.Println("s1: ", s1)
	// fmt.Printf("pointer type: %T\n", pointer_s)
	// fmt.Println("s and pointer", s, pointer_s)
}
