package main

import (
	"fmt"
)

func main ()  {

	ch := make(chan string)

	for i := 0; i< 5000; i++{
		go printHelloword(i, ch)
	}
	for {
		msg := <-ch
		fmt.Println(msg)
	}
	// time.Sleep(time.Millisecond)
}

func printHelloword(i int, ch chan string)  {
	for {
		ch <- fmt.Sprintf("Hello World from go routine %d \n", i)
	}
}