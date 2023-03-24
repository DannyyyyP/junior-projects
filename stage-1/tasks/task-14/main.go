package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fibs := []int{0, 1}
	if a < 2 {
		fmt.Printf("%d\n", a)
	} else if a >= 2 && a <= 3 {
		fmt.Printf("%d\n", a+1)
	} else if a == 4 {
		fmt.Printf("-1\n")
	} else if a > 4 {
		for i := 1; ; i++ {
			b := fibs[i] + fibs[i-1]
			fibs = append(fibs, b)
			if b < a {
				continue
			} else {
				break
			}
		}
	}
	for i := 4; a > 4; i++ {
		if a == fibs[i] {
			fmt.Printf("%d\n", i)
			break
		} else if a < fibs[i] && a > fibs[i-1] {
			fmt.Printf("-1\n")
			break
		}
	}
}
