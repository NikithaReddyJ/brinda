package main

 

import "fmt"

 

func main() {
    pow := make([]int, 10)
    for i, _ := range pow {  //for i,_ := range data
                             // for i := range data
        pow[i] = 2 << uint(i) // == 2* 2**i
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
