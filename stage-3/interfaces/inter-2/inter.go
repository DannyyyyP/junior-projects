// Sample Output:
// 1000010011
// Sample Output:
// [      XXXX]
package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Battery struct {
	Battery string
}

func (b Battery) String() string {
	var countX int
	var countZero int
	batteryInX := make([]string, 0, 10)
	for _, v := range b.Battery {
		if string(v) == "1" {
			countX++
		}
	}
	for _, v := range b.Battery {
		if string(v) == "0" {
			countZero++
		}
	}
	for i := 0; i < countZero; i++ {
		batteryInX = append(batteryInX, " ")
	}
	for i := 0; i < countX; i++ {
		batteryInX = append(batteryInX, "X")
	}
	var buf bytes.Buffer
	buf.WriteRune('[')
	buf.WriteString(strings.Join(batteryInX, ""))
	buf.WriteRune(']')
	return buf.String()
}

func main() {
	var input string
	fmt.Scan(&input)
	var batteryForTest Battery
	batteryForTest.Battery = input
	fmt.Println(batteryForTest.String())
}
