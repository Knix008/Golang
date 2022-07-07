package main

import (
	"fmt"
)

func f1(out chan<- int64, in <-chan int64) {
	out <- (<-in)
}

func f2(out chan int64, in chan int64) {
	out <- (<-in)
}

func main() {
	out := make(chan int64, 10)
	in := make(chan int64, 10)
	in <- 100
	f1(out, in)
	in <- 200
	f2(out, in)
	fmt.Println(<-out)
	fmt.Println(<-out)
}
