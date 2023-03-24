package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringOriginal string
	fmt.Scan(&stringOriginal)
	stringSplitted := strings.Split(stringOriginal, "")
	stringJoined := strings.Join(stringSplitted, "*")
	fmt.Printf("%s", stringJoined)
}
