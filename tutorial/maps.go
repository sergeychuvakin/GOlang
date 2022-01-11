package main

import (
	"fmt"
)

var list_of_maps = make([]map[string]int, 10) // make should deal with inintial size

func main() {

	//var userData map[string]string
	var userData = make(map[string]int)
	userData["a"] = 1
	userData["b"] = 2
	userData["c"] = 3
	// fmt.Println(userData)
	// fmt.Printf("%T\n", strconv.FormatUint(uint64(10), 10))
	// fmt.Printf("%T\n", string(uint64(10)))
	list_of_maps = append(list_of_maps, userData, userData, userData)
	fmt.Println(list_of_maps)
}
