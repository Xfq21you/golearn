package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	a, b := 1, 2
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
}