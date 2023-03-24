package main

import (
	"fmt"
)

func divide(a int, b int) (int, error) {
	return a / b, nil
}

func main() {
	var (
		a int
		b int
	)
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		answer, err := divide(a, b)
		if err == nil {
			fmt.Println(answer)
		} else {
			fmt.Println("ошибка")
		}
	}
}
