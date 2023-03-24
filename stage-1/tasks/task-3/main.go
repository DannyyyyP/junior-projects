package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	if k > 0 && k < 86399 {

		hours := k / 3600
		minutes := k % 3600 / 60

		fmt.Printf("It is %v hours %v minutes.", hours, minutes)
	}
}
