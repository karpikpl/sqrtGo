package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := x / 2
	oldZ := float64(0)
	i := 1

	for math.Abs(oldZ-z) > 0.001 {
		oldZ = z
		z = z - (z*z-x)/2*z
		i++
	}

	return z, i
}

func Sqrt2(x float64, z float64, i int) (float64, int) {
	i++
	if newZ := z - (z*z-x)/2.0*z; math.Abs(newZ-z) < 0.001 {
		return newZ, i
	} else {
		return Sqrt2(x, newZ, i)
	}
}

func main() {

	result, iterations := Sqrt(3)
	result2, iterations2 := Sqrt2(3, 3/2.0, 0)

	fmt.Println("Result is", result, "after", iterations, "iterations and math.Sqrt is", math.Sqrt(3))
	fmt.Println("Result of Sqrt(3) is", result2, "after", iterations2)
}
