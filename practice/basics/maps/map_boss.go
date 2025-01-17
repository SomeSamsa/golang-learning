package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	line := strings.Fields(s)
	for v := range line {
		m[line[v]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
