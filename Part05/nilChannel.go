package main

func main() {
	var c chan string // The channel is initialized to "nil".
	close(c)
}
