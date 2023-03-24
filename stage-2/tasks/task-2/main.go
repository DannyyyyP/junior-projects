package main

import (
	"fmt"
	"math"
)

func gipo (a,b float64) float64 {
	c:= math.Sqrt(a*a +b*b)
	return c
}

func main() {
	var (
		a,b float64
	)
	fmt.Scan(&a,&b)
	fmt.Printf("%.f\n",gipo(a,b))
}