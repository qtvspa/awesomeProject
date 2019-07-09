package main

import "fmt"

func main() {
	ch := make(chan int, 2222)
	ch <- 1
	ch <- 432
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

