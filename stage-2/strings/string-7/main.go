package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	textRunes := []rune(text)
	textInv := []rune(text)
	for i := 0; i < len(textRunes); i++ {
		textInv[i] = textRunes[(len(textRunes)-1)-i]
	}
	if strings.EqualFold(string(textRunes), string(textInv)) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
