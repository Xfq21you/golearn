package main

import (
	"fmt"
)

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a, b int = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
