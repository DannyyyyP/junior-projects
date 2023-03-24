package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	sByte := []byte(s)
	var max, d int
	var err error
	for _, digit := range sByte {
		if d, err = strconv.Atoi(string(digit)); err != nil {
			log.Fatal(err)
		}
		if d > max {
			max = d
		}
	}
	fmt.Println(max)
}

// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

// Выходные данные

// Выведите максимальную цифру, которая встречается во введенной строке.

// Sample Input:

// 1112221112
// Sample Output:

// 2
