package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := [] int{3, 4, 5, 6, 72, 10}
	sort.Ints(a)
	for i,v := range a{
		fmt.Println(i, v)
	}
}
