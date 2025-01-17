package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d"}
	fmt.Println(arr)

	a := arr[1:3]
	b := arr[0:2]

	fmt.Println(a, b)

	b[0] = "some"
	fmt.Println(a, b)

}
