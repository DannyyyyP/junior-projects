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
  arrCopy := workArray
  for i, v := range change {
    if i%2 == 0 {
      workArray[v] = arrCopy[i+1]
    } else {
      workArray[v] = arrCopy[i-1]
    }
  }

  for _, v := range workArray {
    fmt.Print(v, " ")
  }
}


for i, v := range change {
    if i%2 == 0 {
      workArray[v] = arrCopy[i+1]
      continue
    }
    workArray[v] = arrCopy[i-1]
  }