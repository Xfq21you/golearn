package main

import "fmt"

func main() {
	//x := 0

	// if x > 10        // Error: missing condition in if statement
	// {
	// }
	/*
		if n := "abc"; x > 0 { // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
			println(n[2])
		} else if x < 0 { // 注意 else if 和 else 左大括号位置。
			println(n[1])
		} else {
			println(n[0])
		}
	*/
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
