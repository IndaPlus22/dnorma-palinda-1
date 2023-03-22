package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		b := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)

		if math.Abs(z-b) < 0.000001 {
			return z
		}
	}
}
func main() {
	fmt.Println(Sqrt(2))
}
