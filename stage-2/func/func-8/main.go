package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var nFibo = fibonacci(n)
	fmt.Println(nFibo)
}

func fibonacci(n int) int {
	fibo := []int{1, 1}
	for i := 3; i <= n && n >= 2; i++ {
		fibo = append(fibo, fibo[0]+fibo[1])
		copy(fibo[0:], fibo[1:])
		fibo[2] = 0
		fibo = fibo[:2]
	}
	return fibo[1]
}
