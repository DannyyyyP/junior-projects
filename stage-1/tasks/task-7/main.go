package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int
	var counter int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			counter++
		}
	}
	fmt.Println(counter)

}
