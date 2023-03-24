package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// var n int = 35567
	var n int
	fmt.Scan(&n)
	var sum int
	for {
		m := strconv.Itoa(n)
		for i := 0; i < len(m); i++ {
			p, err := strconv.Atoi(string(m[i]))
			if err != nil {
				log.Fatal(err)
			}
			sum += p
		}
		n = sum
		sum = 0
		if n < 10 {
			break
		}
	}
	fmt.Printf("%d\n", n)
}
