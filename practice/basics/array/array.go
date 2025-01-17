package main

import "fmt"

func main() {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	num := [6]string{"1, 2, 3, 4, 5, 6"}
	fmt.Println(num[7])
}
