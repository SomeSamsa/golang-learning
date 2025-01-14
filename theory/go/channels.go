package main

import (
	"fmt"
)

func sendData(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for data := range ch {
		fmt.Println("Data Received:", data)
	}
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	receiveData(ch)
}
