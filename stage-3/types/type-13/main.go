package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	textBySplit := strings.Split(text, " ")
	s1, s2 := textBySplit[0], textBySplit[1]
	fmt.Println(adding(s1, s2))
}

func adding(s1, s2 string) int64 {
	var (
		textSlice1 = []byte(s1)
		textSlice2 = []byte(s2)
		textNum    []byte
		num1       int
		num2       int
		sumOfNum   int
		err        error
	)
	for _, v := range textSlice1 {
		if v >= 48 && v <= 57 {
			textNum = append(textNum, v)
			num1, err = strconv.Atoi(string(textNum))
			if err != nil {
				panic(err)
			}
		}
	}
	textNum = textNum[:0]
	for _, v := range textSlice2 {
		if v >= 48 && v <= 57 {
			textNum = append(textNum, v)
			num2, err = strconv.Atoi(string(textNum))
			if err != nil {
				panic(err)
			}
		}
	}
	sumOfNum = num1 + num2
	return int64(sumOfNum)
}

// %^80 hhhhh20&&&&nd
