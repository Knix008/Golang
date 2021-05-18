package main

import "fmt"
import "runtime"

type twoInts struct {
	X int64
	Y int64
}

func main() {
	fmt.Printf("Hello, World!!!")
	fmt.Printf("\nGOOS:%s", runtime.GOOS)
	fmt.Printf("\n%s", runtime.GOARCH)
	fmt.Printf("\nGOROOT:%s", runtime.GOROOT())
	fmt.Printf("\nThis is only for test!!!")
	fmt.Printf("\nNo. of CPU:%d", runtime.NumCPU())
	fmt.Printf("\nCompiler:%s\n", runtime.Compiler)

	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}

func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}
