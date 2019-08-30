package main

import "fmt"

func add(x int, y float32) float32 {
	return float32(x) + y
	}
	func sub(x int, y float32) float32 {
	return float32(x) - y
	}
	func mul(x int, y float32) float32 {
	return float32(x) * y
	}
	func div(x int, y float32) int {
	return x / int(y)
	}

func main() {
	fmt.Println(add(42, 13.2))
	fmt.Println(sub(32, 13.2))
	fmt.Println(mul(42, 13.2))
	fmt.Println(div(42, 12.2))
}
