package main

import (
	"fmt"
	"math"
)

var c, python, java bool

//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1 << 64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)


const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int)int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	var i int
	var f float64
	var b bool
	var s string
	k := 333
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) wrong
	fmt.Println(add(33,88))
	fmt.Println(swap("world", "hello"))
	fmt.Println(split(16))
	fmt.Println(i, c, python, java, k)

	//const f = "%T(%v)\n"
	//fmt.Printf(f, ToBe, ToBe)
	//fmt.Printf(f, MaxInt, MaxInt)
	//fmt.Printf(f, z, z)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var xx, yy int = 3, 4
	var ff float64 = math.Sqrt(float64(xx*xx + yy*yy))
	var zz int = int(ff)
	fmt.Println(xx, yy, zz)

	v := 42.2423
	fmt.Printf("v is of type %T\n", v)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int)(x ,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
