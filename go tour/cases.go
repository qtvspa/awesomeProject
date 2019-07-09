package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n ,lim float64) float64{
	// if语句可以在条件之前执行一个简单的语句
	if v := math.Pow(x, n); v<lim{
		return v
	}else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main(){
	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os{
	// runtime.GOOS用于返回当前的操作系统
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Thursday?")
	today := time.Now().Weekday()
	switch time.Thursday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
		fallthrough
	case today + 2:
		fmt.Println("In two days.")
		// fallthrough
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}

	//// defer 语句会延迟函数的执行直到上层函数返回
	//defer fmt.Println("world")
	//fmt.Println("hello")
	//// 延迟的函数调用会被压入一个栈中 当函数返回时 会按照后进先出顺序被延迟函数调用
	//fmt.Println("counting")
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("done")
}