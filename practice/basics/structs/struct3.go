package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{3, 4}
	p := &v
	p.X = 78
	fmt.Println(v)
}
