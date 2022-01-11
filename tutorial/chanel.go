package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int, 2)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	// x, y := <-ch, <-ch // receive from c
	// fmt.Println(x, y)

	// fmt.Println(<-ch, <-ch)
	// fmt.Println(cap(ch))

	for i := 1; i <= cap(ch); i++ {
		fmt.Println(<-ch)
	}
}
