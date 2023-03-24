package main

import (
	"fmt"
)

func main() {
	var a int
	m := make(map[int]int)
	for i := 1; i <= 10; i++ {
		fmt.Scan(&a)
		if _, ok := m[i]; !ok {
			m[i] = work(a)
			fmt.Printf("%d ", m[i])
		} else {
			fmt.Printf("%d ", m[i])
		}
	}
}
