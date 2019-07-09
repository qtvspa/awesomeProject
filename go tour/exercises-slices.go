package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	canvas := make([][]uint8, dx)
	for i:=0; i<dx; i++ {
		canvas[i] = make([]uint8, dy)
		for j:=0; j<dy; j++ {
			canvas[i][j] = uint8(i*j)
		}
	}
	return canvas
}

func main() {
	pic.Show(Pic)
}