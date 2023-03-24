package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(i uint) uint {
		iToString := strconv.Itoa(int(i))
		var iWithout []rune
		for _, v := range iToString {
			if v%2 == 0 && v != 48 {
				iWithout = append(iWithout, v)
			}
		}
		if string(iWithout) == "" {
			return 100
		} else {
			g, err := strconv.Atoi(string(iWithout))
			if err != nil {
				panic(err)
			} else if g == 0 {
				return 100
			} else {
				return uint(g)
			}
		}
	}
	var a uint
	fmt.Scan(&a)
	fmt.Println(fn(a))
}
