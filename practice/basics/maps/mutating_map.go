package main

import "fmt"

func print(x int) {
	fmt.Printf("The value: %v", x)
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 48
	print(m["Answer"])

	m["Answer"] = 42
	print(m["Answer"])

	delete(m, "Answer")
	print(m["Answer"])

	v, ok := m["Answer"]
	print(v)
	fmt.Println("Present? ", ok)
}
