package main

import (
	"fmt"
)

func main() {
	a, b := sumInt(0, 1)
	fmt.Printf("%d, %d\n", a, b)
}

func sumInt(args ...int) (int, int) {
	var summa []int
	var b int
	for _, elem := range args {
		summa = append(summa, elem)
	}
	for _, v := range summa {
		b += v
	}
	a := len(summa)
	return a, b
}
