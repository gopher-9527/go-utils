package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{}
	for i := 0; i < 10; i++ {
		pool.Put(&Connection{})
	}
	conn := pool.Get().(*Connection)
	fmt.Println("before gc: ", *conn)

	runtime.GC()
	runtime.GC()
	conn = pool.Get().(*Connection)
	fmt.Println("after gc: ", *conn)
}

func condDemo() {
	// 定义一个条件变量
	var cond = sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			// 加锁
			cond.L.Lock()
			ready++
			// 执行任务
			fmt.Println("goroutine", i, "running...")
			// 解锁
			cond.L.Unlock()
			cond.Broadcast()
		}(i)
	}

	cond.L.Lock()
	for ready != 10 {
		cond.Wait()
		println("ready:", ready)
	}
	cond.L.Unlock()
	println("done")
}

type Connection struct{}

func poolDemo() sync.Pool {
	pool := sync.Pool{}
	pool.New = func() interface{} {
		return &Connection{}
	}
	return pool
}
