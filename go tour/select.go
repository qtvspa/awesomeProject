package main

import (
	"fmt"
)
// select 语句使一个 Go 程可以等待多个通信操作。
// select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

// 当 select 中的其它分支都没有准备好时，default 分支就会执行。

//func main() {
//	tick := time.Tick(100 * time.Millisecond)
//	boom := time.After(500 * time.Millisecond)
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick.")
//		case <-boom:
//			fmt.Println("BOOM!")
//			return
//		default:
//			fmt.Println("    .")
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
//}