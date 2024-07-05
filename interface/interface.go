package main

import "fmt"

type Sayer interface {
	say()
}

type cat struct{}
type dog struct{}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (d dog) say() {
	fmt.Println("汪汪汪")
}

func main() {
	var x Sayer
	c := cat{}
	d := dog{}
	x = c
	x.say()
	x = d
	x.say()
}
