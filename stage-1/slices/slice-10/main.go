package main

import "fmt"

func main() {
	var n int
	var number int
	var array []int
	fmt.Scan(&number)
	if number > 3 {
		for i := 0; i < number; i++ {
			fmt.Scan(&n)
			array = append(array, n)
		}
		fmt.Printf("%d\n", array[3])
	}
}
