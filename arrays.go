package main

 

import "fmt"

 

func main() {
    var a [3]string
    a[0] = "Hello"
    a[1] = "World"
    a[2] = "2"
    fmt.Println(a[0]+a[1])
    for i, v:= range a{
        fmt.Println(i, v)
    }
    fmt.Println(a)

 

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
