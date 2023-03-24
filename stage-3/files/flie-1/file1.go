package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum int
	in := os.Stdin
	out := os.Stdout
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	if _, err := io.WriteString(out, strconv.Itoa(sum)+"\n"); err != nil {
		panic(err)
	}
}
