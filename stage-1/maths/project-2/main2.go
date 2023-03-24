package main

import "fmt"

func main() {
	var numOfSums int
	fmt.Scan(&numOfSums)

	var sum int
	for numOfSums > 0 {
		var a, b int
		fmt.Scan(&a, &b)
		sum += a + b
		numOfSums--
	}

	fmt.Println(sum)
}