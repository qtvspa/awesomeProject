package main

import "fmt"

////结构体
//type Vertex struct{
//	X int
//	Y int
//}
//
//func main(){
//
//	fmt.Println(Vertex{3, 4})
//	v := Vertex{66, 77}
//	fmt.Println(v.X)
//
//	p := &v
//	p.X = 1e9
//	fmt.Println(v)
//}
type Vertex struct {
	X, Y int
}
type Vertex2 struct {
	Lat, Long float64
}

//var (
//	v1 = Vertex{1, 2}  // 类型为 Vertex
//	v2 = Vertex{X: 1}  // Y:0 被省略
//	v3 = Vertex{}      // X:0 和 Y:0
//	p  = &Vertex{1, 2} // 类型为 *Vertex
//)

//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m map[string]Vertex2

func main() {
	//fmt.Println(v1, p, v2, v3)
	//数组
	//var a[2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//切片
	//pp := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("pp ==", pp)
	//fmt.Println(pp[:4])

	//for i := 0; i < len(pp); i++ {
	//	fmt.Printf("pp[%d] == %d\n", i, pp[i])
	//}

	//b := make([]int, 0, 5)
	//fmt.Println(b, len(b), cap(b))
	//b = append(b, 7, 7, 7)
	//fmt.Println(b)

	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}