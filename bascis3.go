package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a bool       = false
	b uint64     = 1<<64 - 1
	c complex128 = cmplx.Sqrt(-5 + 12i)
	d complex128 = 20 + 9i
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
}
