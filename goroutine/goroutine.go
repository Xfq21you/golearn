package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	func hello() {
		fmt.Println("Hello Goroutine!")
	}

	func main() {
		go hello()
		fmt.Println("main goroutine done!")
		time.Sleep(time.Second)
	}

var wg sync.WaitGroup

	func hello(i int) {
		defer wg.Done()
		fmt.Println("Hello Goroutine!", i)
	}

	func main() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go hello(i)
		}
		wg.Wait()
	}

	func main() {
		go func() {
			i := 0
			for {
				i++
				fmt.Printf("new goroutine: i = %d\n", i)
				time.Sleep(time.Second)
			}
		}()
		i := 0
		for {
			i++
			fmt.Printf("main goroutine: i = %d\n", i)
			time.Sleep(time.Second)
			if i == 2 {
				break
			}
		}
	}

	func main() {
		go func(s string) {
			for i := 0; i < 2; i++ {
				fmt.Println(s)
			}
		}("world")
		for i := 0; i < 2; i++ {
			runtime.Gosched()
			fmt.Println("hello")
		}
	}

	func main() {
		go func() {
			defer fmt.Println("A.defer")
			func() {
				defer fmt.Println("B.defer")
				runtime.Goexit()
				defer fmt.Println("C.defer")
				fmt.Println("B")
			}()
			fmt.Println("A")
		}()
		for {

		}
	}
*/
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
