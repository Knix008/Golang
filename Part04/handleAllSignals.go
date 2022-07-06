package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handle(sig)
			case syscall.SIGTERM:
				handle(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall.SIGUSR2!")
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()
	i := 3
	for {
		fmt.Printf(".")
		time.Sleep(3 * time.Second)
		if i == 3 {
			sigs <- syscall.SIGUSR2
		} else if i == 2 {
			sigs <- syscall.SIGUSR1
		} else {
			sigs <- syscall.SIGTERM
		}
		i--
	}
}
