package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	fmt.Scan(&a)
	var c int
	fmt.Scan(&c)
	except := strconv.Itoa(c)
	b := strconv.Itoa(a)
	var number []string
	for _, e := range b {
		if string(e) != except {
			number = append(number, string(e))
		} else if string(e) == except {
			continue
		}
	}
	for i := 0; i < len(number); i++ {
		fmt.Printf("%s", number[i])
	}
}