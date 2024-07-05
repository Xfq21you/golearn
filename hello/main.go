package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 

func main(){  // main函数，是程序执行的入口
    fmt.Println("Hello World!")  // 在终端打印 Hello World!

    var (
        a string
        b int
        c bool
    )
    const (
        pi = 3.14
        e = 2.7
        d = 12
        f
    )
    const (
        n1 = iota
        n2
        _
        n4
    )
    s1 := `asudhiua
    skfd
    sfds
    `
    fmt.Println(a, b, c, d, e, f, pi, n1, n4, s1)
}

/*
func foo() (int, string) {
    return 10, "Q1mi"
}
func main() {
    x, _ := foo()
    _, y := foo()
    fmt.Println("x=", x)
    fmt.Println("y=", y)
}
*/