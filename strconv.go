package main

import (
    "fmt"
    "strconv"
)

func main() {

i := strconv.Itoa('A')
j,_ := strconv.Atoi("42")
fmt.Println(i,j)
fmt.Printf("%T %T", i,j)
}
