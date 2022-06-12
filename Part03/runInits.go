package main

import (
	"fmt"
	"initA"
	"initB"
)

func init() {
	fmt.Println("init() runInits")
}
func main() {
	initA.FromA()
	initB.FromB()
}
