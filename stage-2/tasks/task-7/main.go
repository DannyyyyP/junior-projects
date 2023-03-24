package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		p, v, k float64
	)
	fmt.Scan(&p, &v, &k)
	M(p, v)
}

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	w := math.Sqrt(k / M())
	return w
}

func T() float64 {
	t := 6 / W()
	return t
}
