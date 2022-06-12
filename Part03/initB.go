package initB

import (
	"fmt"
	"initA"
)

func init() {
	fmt.Println("init() b")
}
func FromB() {
	fmt.Println("fromB()")
	initA.FromA()
}
