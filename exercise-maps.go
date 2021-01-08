package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var counter make([string]int, 0)
	words := strings.Fields(s)
	for _, w in range words {
		if v, ok := counter[w]; ok {
			counter[w] += 1
		} else {
			counter[w] = 0
		}
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
