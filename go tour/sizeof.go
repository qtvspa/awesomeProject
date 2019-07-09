package main

import (
	"fmt"
	"unsafe"
)

func main()  {
 	var b string = "awerwerwererwwrer"
 	fmt.Println(unsafe.Sizeof(b))
 	fmt.Println(len(b))
}
