package tutorial

import (
	"fmt"
	"strings"
)

func main() {

	var NameS []string
	var Ns = []uint{}

	for {

		var Name string
		var Name2 string
		var N uint

		fmt.Print("Your first Name: ")
		fmt.Scanln(&Name)

		fmt.Print("Your second Name: ")
		fmt.Scanln(&Name2)

		fmt.Print("Number of tickets: ")
		fmt.Scanln(&N)

		NameS = append(NameS, Name+" "+Name2)
		Ns = append(Ns, N)

		//fmt.Println(bookingS)

		// for-each

		// fmt.Println(end)
		fmt.Println(len(NameS))
		for _, i := range NameS {

			var names = strings.Fields(i)
			var first = names[0]
			//fmt.Println(names)
			fmt.Println("Your first name:", first)

			//fmt.Println(id, i)
		}

		//end := len(NameS) == 2
		//var end bool = len(NameS) > 2
		if len(NameS) >= 1 {
			break
		} else if true {
			continue
		}
	}

}
