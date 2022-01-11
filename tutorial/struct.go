package main

import "fmt"

type UserData struct {
	a string
	b string
	c uint
	d bool
	g []string
}

var list = make([]UserData, 0)

func main() {

	var user = UserData{
		a: "a",
		b: "b",
		c: 1,
		d: true,
		g: []string{"1", "s"},
	}

	list = append(list, user)
	fmt.Println(list)
}
