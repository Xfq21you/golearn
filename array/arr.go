package main

import "fmt"
/*
var arr0[5] int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}
*/

var arr0[5][3] int
var arr1[2][3] int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	//a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	//fmt.Println(arr0, arr1)
	//fmt.Println(a, b)
	for k1, v1 := range arr1 {
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d) = %d ", k1, k2, v2)
		}
		fmt.Println()
	}
}
/*
package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1[5] int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}
*/