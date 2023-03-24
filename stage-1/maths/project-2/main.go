package main

import "fmt"

func main() {
	var numOfSums int
	fmt.Scan(&numOfSums)

	var sum int
	var flag bool = true
	for numOfSums > 0 && flag {
		var a, b int
		fmt.Scan(&a, &b)
		sum += a + b
		numOfSums--
		if sum > 10 {
			flag = false
		}
	}

	fmt.Println(sum)
}
