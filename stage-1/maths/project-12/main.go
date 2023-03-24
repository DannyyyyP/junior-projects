package main

import (
	"fmt"
)

func main() {

	workArray := [10]uint8{}
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	change := [6]uint8{}
	for i := 0; i < len(change); i++ {
		fmt.Scan(&change[i])
	}
	/*workArraynew := workArray

	workArray[change[0]] = workArraynew[change[1]]
	workArray[change[1]] = workArraynew[change[0]]
	workArray[change[2]] = workArraynew[change[3]]
	workArray[change[3]] = workArraynew[change[2]]
	workArray[change[4]] = workArraynew[change[5]]
	workArray[change[5]] = workArraynew[change[4]]
	*/

	for i := 0; i < 3; i++ {
		workArray[change[2*i]], workArray[change[2*i+1]] = workArray[change[2*i+1]], workArray[change[2*i]]
	}
		
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
