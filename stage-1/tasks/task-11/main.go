package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for {
		if b%7 == 0 {
			break
		} else if b%7 != 0 {
			b = b - 1
		} else if b < a {
			break
		}
	}
	if b < a {
		fmt.Println("NO")
	} else {
		fmt.Println(b)
	}
}
