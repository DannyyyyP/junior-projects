package main

import (
	"fmt"
)

func main() {

	var a int

	fmt.Scan(&a)

	if a > 0 && a < 1000 {
		a1 := a / 100
		a2 := (a / 10) % 10
		a3 := a % 10
		asum := a1 + a2 + a3

		fmt.Println(asum)
	}
}
