package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v = Vertex{1, 2}
	s = Vertex{}
	r = Vertex{X: 3}
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v, s, r, *p)
}
