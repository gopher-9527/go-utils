package main

import (
	"fmt"
	"runtime"
	"sync"
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

type Connection struct{}
