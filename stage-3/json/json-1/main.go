package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int
}
type Group struct {
	Students []Student
}
type Rating struct {
	Average float32
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var group Group
	if err = json.Unmarshal(input, &group); err != nil {
		panic(err)
	}
	studentCount := len(group.Students)
	var ratingCount int
	for _, std := range group.Students {
		ratingCount += len(std.Rating)
	}
	var rating Rating
	rating.Average = float32(ratingCount) / float32(studentCount)

	data, err := json.MarshalIndent(rating, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(data))
}
