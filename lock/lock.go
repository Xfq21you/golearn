package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

	func add() {
		for i := 0; i < 5000; i++ {
			lock.Lock()
			x = x + 1
			lock.Unlock()
		}
		wg.Done()
	}

	func main() {
		wg.Add(2)
		go add()
		go add()
		wg.Wait()
		fmt.Println(x)
	}

var (

	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex

)

	func write() {
		rwlock.Lock()
		x = x + 1
		time.Sleep(time.Millisecond)
		rwlock.Unlock()
		wg.Done()
	}

	func read() {
		rwlock.Lock()
		time.Sleep(time.Millisecond)
		rwlock.Unlock()
		wg.Done()
	}

	func main() {
		start := time.Now()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go write()
		}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go read()
		}
		wg.Wait()
		end := time.Now()
		fmt.Println(end.Sub(start))
	}

var m = sync.Map{}

	func main() {
	    wg := sync.WaitGroup{}
	    for i := 0; i < 20; i++ {
	        wg.Add(1)
	        go func(n int) {
	            key := strconv.Itoa(n)
	            m.Store(key, n)
	            value, _ := m.Load(key)
	            fmt.Printf("k=:%v,v:=%v\n", key, value)
	            wg.Done()
	        }(i)
	    }
	    wg.Wait()
	}
*/
var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go add() // 普通版add函数 不是并发安全的
		// go mutexAdd() // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
