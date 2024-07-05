package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	func main() {
		// fmt.Print("在终端打印该信息。")
		// name := "枯藤"
		// fmt.Printf("我是：%s\n", name)
		// fmt.Println("在终端打印单独一行显示")

		// 向标准输出写入内容
		// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
		// fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		// if err != nil {
		// 	fmt.Println("打开文件出错，err:", err)
		// 	return
		// }
		// name := "枯藤"
		// // 向打开的文件句柄中写入内容
		// fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

		// s1 := fmt.Sprint("枯藤")
		// name := "枯藤"
		// age := 18
		// s2 := fmt.Sprintf("name:%s,age:%d", name, age)
		// s3 := fmt.Sprintln("枯藤")
		// fmt.Println(s1, s2, s3)

		// fmt.Printf("%v\n", 100)
		// fmt.Printf("%v\n", false)
		// o := struct{ name string }{"枯藤"}
		// fmt.Printf("%v\n", o)
		// fmt.Printf("%#v\n", o)
		// fmt.Printf("%T\n", o)
		// fmt.Printf("100%%\n")

		// n := 65
		// fmt.Printf("%b\n", n)
		// fmt.Printf("%c\n", n)
		// fmt.Printf("%d\n", n)
		// fmt.Printf("%o\n", n)
		// fmt.Printf("%x\n", n)
		// fmt.Printf("%X\n", n)

		// f := 12.34
		// fmt.Printf("%b\n", f)
		// fmt.Printf("%e\n", f)
		// fmt.Printf("%E\n", f)
		// fmt.Printf("%f\n", f)
		// fmt.Printf("%g\n", f)
		// fmt.Printf("%G\n", f)

		// s := "枯藤"
		// fmt.Printf("%s\n", s)
		// fmt.Printf("%q\n", s)
		// fmt.Printf("%x\n", s)
		// fmt.Printf("%X\n", s)

		// a := 18
		// fmt.Printf("%p\n", &a)
		// fmt.Printf("%#p\n", &a)

		// n := 88.88
		// fmt.Printf("%f\n", n)
		// fmt.Printf("%9f\n", n)
		// fmt.Printf("%.2f\n", n)
		// fmt.Printf("%9.2f\n", n)
		// fmt.Printf("%9.f\n", n)

		s := "枯藤"
		fmt.Printf("%s\n", s)
		fmt.Printf("%5s\n", s)
		fmt.Printf("%-5s\n", s)
		fmt.Printf("%5.7s\n", s)
		fmt.Printf("%-5.7s\n", s)
		fmt.Printf("%5.2s\n", s)
		fmt.Printf("%05s\n", s)
	}

	func main() {
		var (
			name    string
			age     int
			married bool
		)
		//fmt.Scan(&name, &age, &married)
		//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

		//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
		//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

		fmt.Scanln(&name, &age, &married)
		fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	}
*/
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	bufioDemo()
}
