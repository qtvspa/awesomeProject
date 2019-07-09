package main

import (
	"fmt"
	"math"
)

//func main() {
//	//names := [4]string{
//	//	"John",
//	//	"Paul",
//	//	"George",
//	//	"Ringo",
//	//}
//	//fmt.Println(names)
//	//
//	//a := names[0:2]
//	//b := names[1:3]
//	//fmt.Println(a, b)
//	//// 更改切片的元素会修改其底层数组中对应的元素
//	//b[0] = "XXX"
//	//fmt.Println(a, b)
//	//fmt.Println(names)
//	// 创建一个井字板（经典游戏）
//	//board := [][]string{
//	//	[]string{"_", "_", "_"},
//	//	[]string{"_", "_", "_"},
//	//	[]string{"_", "_", "_"},
//	//}
//
//	//// 两个玩家轮流打上 X 和 O
//	//board[0][0] = "X"
//	//board[2][2] = "O"
//	//board[1][2] = "X"
//	//board[1][0] = "O"
//	//board[0][2] = "X"
//	//
//	//for i := 0; i < len(board); i++ {
//	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
//	//}
//}


type Vertex4 struct{
	X, Y float64
}

func (v Vertex4) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex4) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	v := Vertex4{3, 4}
	//fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}