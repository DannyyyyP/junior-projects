package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	textBySplit := strings.Split(text, ";")
	s1, s2 := textBySplit[0], textBySplit[1]
	fmt.Printf("%.4f", division(s1, s2))
}
func division(s1, s2 string) float64 {
	var (
		textSlice1 = []byte(s1)
		textSlice2 = []byte(s2)
		textNum    []byte
		num1       float64
		num2       float64
		sumOfNum   float64
		err        error
	)
	for _, v := range textSlice1 {
		if v >= 48 && v <= 57 {
			textNum = append(textNum, v)
		}
		if v == 44 {
			textNum = append(textNum, byte('.'))
		}
	}
	num1, err = strconv.ParseFloat(string(textNum), 4)
	if err != nil {
		panic(err)
	}
	textNum = textNum[:0]
	for _, v := range textSlice2 {
		if v >= 48 && v <= 57 {
			textNum = append(textNum, v)
		}
		if v == 44 {
			textNum = append(textNum, byte('.'))
		}
	}
	num2, err = strconv.ParseFloat(string(textNum), -1)
	if err != nil {
		panic(err)
	}
	sumOfNum = num1 / num2
	return float64(sumOfNum)
}

// 1 149,6088607594936;1 179,0666666666666
