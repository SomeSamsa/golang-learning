package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printMessage("Hello from Goroutine")
	fmt.Println("Main function")
	time.Sleep(11 * time.Second)
}
