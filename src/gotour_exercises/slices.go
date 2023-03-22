package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			x[j] = uint8(i*i + j*j)
		}
		s[i] = x
	}
	return s
}

func main() {
	pic.Show(Pic)
}
