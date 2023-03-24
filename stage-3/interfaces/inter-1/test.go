package main

import (
	"fmt"
)

func res() {
	var (
		a,b,c float64
	)
	fmt.Scan(&a,&b,&c)
	result:= a + c + b
	fmt.Println(result)
}