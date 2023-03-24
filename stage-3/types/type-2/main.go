package main

import "fmt"

func main() {
	var a int64

	fmt.Scan(&a)
	fmt.Println(convert(a))
}
func convert(a int64) uint16 {
	aLittle := uint16(a)
	return aLittle
}
