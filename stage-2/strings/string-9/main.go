package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	textRune := []rune(text)
	var textChanged []rune
	for i := 0; i < len(textRune); i++ {
		if i%2 != 0 {
			textChanged = append(textChanged, textRune[i])
		}
	}
	fmt.Println(string(textChanged))
}
