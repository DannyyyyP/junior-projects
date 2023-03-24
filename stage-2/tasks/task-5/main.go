package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	var final []string
	for _, d:= range n {
		num, err:= strconv.Atoi(string(d))
		if err != nil {
			log.Fatal(err)
		}
		final = append(final, strconv.Itoa((num * num)))
	}
	fmt.Println(strings.Join(final,""))
}


// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число. 

// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

// Sample Input:

// 9119
// Sample Output:

// 811181
