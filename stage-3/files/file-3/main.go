package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	fs, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buff := make([]byte, fs.Size())
	if _, err = rd.Read(buff); err != nil {
		panic(err)
	}
	str := strings.Split(string(buff), ";")
	for i, st := range str {
		if st == "0" {
			fmt.Println(i + 1)
		}
	}
}
