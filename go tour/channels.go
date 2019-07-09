package main

import "fmt"

// 信道

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{3, 4, 6, 2, -5, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x+y)
}
