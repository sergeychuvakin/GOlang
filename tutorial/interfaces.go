package main

import "fmt"

type animal interface {
	makeSound() // genereic type for objects that have that method.
}

type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("meow")
}
func (d dog) makeSound() {
	fmt.Println("bark!")
}

func main() {

	// c := cat{}
	// d := dog{}
	var c, d animal = &cat{}, dog{}
	d.makeSound()
	c.makeSound()
}
