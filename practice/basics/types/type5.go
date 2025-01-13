package main

import (
	"fmt"
	"math/big"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) - overflow
	bigNum := new(big.Int).Lsh(big.NewInt(1), 100)
	//big.Int type that can hold in it mush larger number than in int and here I am calling instance of this
	//big.NewInt(1) creates the instance of this type with value 1
	//Lsh stand for <<(left shift)
	fmt.Println(bigNum)
}
