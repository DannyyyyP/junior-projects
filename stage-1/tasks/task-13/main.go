package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a []int
	for i := 0; i <= n; i++ {
		var b = math.Pow(2, float64(i))
		if int(b) > n {
			break
		}
		a = append(a, int(b))
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}
