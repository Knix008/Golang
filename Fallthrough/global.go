package main

import "fmt"

var global = 1000

func access() {
	local := 1
	fmt.Println(global)
	fmt.Println(local)
}

func main() {
	var local = 100

	fmt.Println(global)
	fmt.Println(local)

	access()
}
