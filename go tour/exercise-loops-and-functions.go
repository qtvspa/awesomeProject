package main
// 牛顿法实现开方函数
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0 // z = float64(1)
	temp := 0.0 // 用于存放临时的z值
	i := 1

	for{
		z = z - (z * z -x)/(2 * z)
		fmt.Printf("第%d次计算结果为: ", i)
		fmt.Println(z)
		if math.Abs(z - temp) < 0.0000000000001{
			break
		}else{
			temp = z
			i ++
		}
	}
	return z
}

func main(){
	fmt.Println("牛顿法计算结果:",Sqrt(2))
	fmt.Println("math.Sqrt(2)结果", math.Sqrt(2))
}


//func ssqrt(x float64) float64 {
//	z := float64(1)
//	temp := 0.0
//	for{
//		z = z - (z * z - x)/(2 * x)
//		if math.Abs(z-temp) < 0.00001{
//			break
//		}else {
//			temp = z
//		}
//	}
//	return z
//}

















