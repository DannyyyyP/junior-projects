// fdsghdfgjsdDD1

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)
	passwordRune := []rune(password)
	n := 1

	if len(passwordRune) >= 5 {
		for i := 0; i < len(passwordRune); i++ {
			if unicode.Is(unicode.Latin, passwordRune[i]) || unicode.IsNumber(passwordRune[i]) {
				continue
			} else {
				n = 0
				break
			}
		}
	} else {
		n = 0
	}
	if n == 1 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
