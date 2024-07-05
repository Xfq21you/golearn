package main

import (
    "fmt"
	"math"
)
/*
func main() {
	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '点'
	fmt.Println(string(runeS2))
}
*/
func main() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}