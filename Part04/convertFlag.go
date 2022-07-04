package main

import (
	"flag"
	"fmt"
)

func main() {
	minusI := flag.Int("i", 100, "i parameter")
	flag.Parse()
	i := *minusI
	fmt.Println(i)
}
