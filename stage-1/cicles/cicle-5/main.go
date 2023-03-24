package main

import "fmt"

func main() {

	var n, sum uint
	fmt.Scan(&n)
	for n > 0 {
		var b uint
		fmt.Scan(&b)
		if b > 9 && b%8 == 0 && b < 100 {
			sum += b
			n--
		} else {
			n--
		}
	}
	fmt.Println(sum)
}
