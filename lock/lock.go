package lock

// 通过共享变量的方式

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0 // 全局共享变量，做到多个协程通信

func count(lock *sync.Mutex) {
	lock.Lock() // 上锁
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock() // 解锁
}

func DoGoLock() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go count(lock)
	}

	//runtime.Gosched()

	for {
		lock.Lock() // 上锁
		c := counter
		lock.Unlock() // 解锁

		runtime.Gosched() // 出让时间片

		if c >= 10 {
			break
		}
	}
}
