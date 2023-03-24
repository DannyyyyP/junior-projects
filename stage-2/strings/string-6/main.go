package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(strings.Replace(text, "\n", "", -1))
	textRunes := []rune(text)

	if unicode.IsUpper(textRunes[0]) && textRunes[len(textRunes)-1] == '.' {
		fmt.Fprintln(os.Stdout, "Right")
	} else {
		fmt.Fprintln(os.Stdout, "Wrong")
	}
}
