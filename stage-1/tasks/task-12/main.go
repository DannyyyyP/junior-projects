package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	b := n / 10
	if b == 0 {
		if n == 1 {
			fmt.Printf("%d korova\n", n)
		} else if 1 < n && n < 5 {
			fmt.Printf("%d korovy\n", n)
		} else if n >= 5 && n < 10 {
			fmt.Printf("%d korov\n", n)
		}
	} else if b == 1 {
		if n%10 >= 0 {
			fmt.Printf("%d korov\n", n)
		}
	} else if b > 1 {
		if n%10 == 0 {
			fmt.Printf("%d korov\n", n)
		} else if n%10 == 1 {
			fmt.Printf("%d korova\n", n)
		} else if n%10 >= 1 && n%10 < 5 {
			fmt.Printf("%d korovy\n", n)
		} else if n%10 >= 5 {
			fmt.Printf("%d korov\n", n)
		}
	}
}
