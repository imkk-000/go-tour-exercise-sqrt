package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for c := 0; c <= 10; c++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 0.0; i <= 100; i++ {
		fmt.Printf("Sqrt(%.0f) = %.6f\n", i, Sqrt(i))
	}
}
