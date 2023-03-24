package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string
	fmt.Scan(&text)
	textRune := []rune(text)
	var textConverted []rune
	for i := 0; i < len(textRune); i++ {
		if strings.Count(string(textRune), string((textRune)[i])) == 1 {
			textConverted = append(textConverted, textRune[i])
		}
	}
	fmt.Println(string(textConverted))
}
