package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	var sl []int = arr[2:4]

	fmt.Println(sl)
}
