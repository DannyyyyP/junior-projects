package main

import "fmt"

func main() {
	var sum, finalsum, procent float32
	var years int
	fmt.Scan(&sum, &procent, &finalsum)
	for years = 1; ; years++ {
		sumafter := sum + sum*(procent/100)
		sum = sumafter
		if sumafter < finalsum {
			continue
		} else {
			break
		}
	}
	fmt.Println(years)
}
