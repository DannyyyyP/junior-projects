package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		text    string
		podText string
	)
	fmt.Scan(&text, &podText)
	fmt.Println(strings.Index(text, podText))

}
