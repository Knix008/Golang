package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)
	// Next line will make a panic message.
	//go A(x, y)
	close(x)
	time.Sleep(3 * time.Second)
}
