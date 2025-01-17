package main

import "fmt"

func main() {
	var s []int

	s = append(s, 0)
	print(s)

	s = append(s, 1)
	print(s)

	s = append(s, 2, 3, 4)
	print(s)
}

func print(s []int) {
	fmt.Printf("len = %d cap = %d %d\n", len(s), cap(s), s)
}
