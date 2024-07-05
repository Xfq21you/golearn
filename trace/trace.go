/*
package main

import (

	"fmt"
	"os"
	"runtime/trace"

)

func main() {

		//创建trace文件
		f, err := os.Create("trace.out")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		//启动trace goroutine
		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()

		//main
		fmt.Println("Hello World")
	}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}
