package main

import "fmt"

func main() {
	var n int
	var array []int
	fmt.Scan(&n)
	if n > 0 {
		for i := 0; i < n; i++ {
			var a int
			fmt.Scan(&a)
			array = append(array, a)
		}
		amin := array[0]
		var counter int = 1
		for i := 1; i < n; i++ {
			if array[i] < amin {
				amin = array[i]
				counter = 1
			} else if array[i] == amin {
				counter++
			}
		}
		fmt.Println(counter)
	}
}
