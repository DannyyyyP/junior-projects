package main

import "fmt"

func main() {

	var n int
	var array []int
	var arrayx2 []int
	fmt.Scan(&n)
	if n <= 100 && n >= 0 {

		for i := 0; i < n; i++ {
			var number int
			fmt.Scan(&number)
			array = append(array, number)
		}
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arrayx2 = append(arrayx2, array[i])
		} else {
			continue
		}
	}
	for _, v := range arrayx2 {
		fmt.Print(v, " ")
	}
}
