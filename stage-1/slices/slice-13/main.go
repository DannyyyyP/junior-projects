package main

import "fmt"

// func main() {

// 	// var n int
// 	// var array []int
// 	// var arrayplus []int
// 	// fmt.Scan(&n)
// 	// if n <= 100 && n >= 0 {

// 	// 	for i := 0; i < n; i++ {
// 	// 		var number int
// 	// 		fmt.Scan(&number)
// 	// 		array = append(array, number)
// 	// 	}
// 	// }
// 	// for _, v := range array {
// 	// 	if v > 0 {
// 	// 		arrayplus = append(arrayplus, v)
// 	// 	}
// 	// }
// 	// fmt.Printf("%d ", len(arrayplus))
// 	}

func main() {
	workArray := []int{99, 151, 137, 71, 117, 187, 20, 93, 187, 67, 1, 2, 3, 5, 7, 8}
	numToSwap := workArray[:10]
	swapPosition := workArray[10:]
	result := make([]int, 10)
	for i := range numToSwap {
		for j, pos := range swapPosition {
			if i%2 == 0 {
				if i == pos {
					result[j] = numToSwap[j+1]
					result[j+1] = numToSwap[j]
					break
				} else {
					result[i] = numToSwap[i]
				}
			}
		}
	}
	fmt.Println(result)
}

// 99 137 151 187 117 71 20 187 93 67
