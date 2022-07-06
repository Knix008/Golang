package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("os.Interrupt Caught!!!:", sig)
				os.Exit(1)
			//case syscall.SIGINFO:
			case syscall.SIGTERM:
				fmt.Println("SIGTERM Caught!!!")
				handleSignal(sig)
				return
			}
		}
	}()
	i := 3
	for {
		fmt.Printf(".")
		time.Sleep(3 * time.Second)
		if i == 0 {
			fmt.Println("Sending signal")
			sigs <- syscall.SIGTERM
			time.Sleep(3 * time.Second)
			break
		} else {
			i--
		}
	}
}
