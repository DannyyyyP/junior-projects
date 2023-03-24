package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)

	var isEqual bool
	for _, symbolA := range aStr {
		for _, symbolB := range bStr {
			if symbolA == symbolB {
				fmt.Printf("a is equal to b, a: %s, b: %s\n", string(symbolA), string(symbolB))
				isEqual = true
			}
		}
	}

	if !isEqual {
		fmt.Println("no one aStr's symbol is equal to bStr")
	}
}
