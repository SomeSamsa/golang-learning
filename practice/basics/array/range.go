package main

import "fmt"

func main() {
	var some = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range some {
		fmt.Printf("2**%d: %d\n", i, v)
	}
}
