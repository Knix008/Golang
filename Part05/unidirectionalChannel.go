package main

import (
	"fmt"
)

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	// Error
	r := <-c
	fmt.Println(r)
}

func main() {
	c := make(chan int, 1)
	d := make(chan int, 1)
	f1(c, 100)
	f2(d, 200)
}
