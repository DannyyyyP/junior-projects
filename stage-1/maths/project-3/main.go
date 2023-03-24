package main

import "fmt"

func main() {

	var a, b uint
	fmt.Scan(&a, &b)
	var sum uint
	for ; a < 100 && b < 100 && a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}
