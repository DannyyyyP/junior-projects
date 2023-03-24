package main

import "fmt"

func main() {
    var n int
    
        for {
            fmt.Scan(&n)
            if n < 10 {
            continue
            } else if n > 100 { 
             break
             } else {
            fmt.Println(n)
            }        
}
}