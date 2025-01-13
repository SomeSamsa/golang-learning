package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		if z == math.Sqrt(x) {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
