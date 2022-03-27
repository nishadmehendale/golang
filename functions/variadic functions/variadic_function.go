package main

import (
	"fmt"
)

func main() {
	bestFinish := getBestFinish(10, 2, 3, 4, 5, 6, 1, 12, 4, 5)
	fmt.Println("Best Finish:", bestFinish)
}

func getBestFinish(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
