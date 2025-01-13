package main

import (
	"fmt"
	"math"
)

func powder(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%v >= %v\n", v, lim)
		return lim
	}
}

func main() {
	fmt.Println(
		powder(3, 2, 10),
		powder(3, 3, 20),
	)
}
