package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Implement the square root function using Newton's method.
// In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating: z - (z*z - x) / (2 * z)
// To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

/*

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:
	z := float64(1)
	z := 1.0

*/

func Sqrt(x float64) float64 {
	fn := calculateSqrt(x)
	y, z := float64(0), float64(1)
	for math.Abs(z-y) > 0.001 {
		y = z
		z = fn()
	}
	return z

}

func calculateSqrt(x float64) func() float64 {

	z := float64(rand.Intn(int(x)))
	return func() float64 {
		z = z - (z*z-x)/2*z
		return z
	}
}

func main() {
	fmt.Println("a")
	fmt.Println(Sqrt(2))
}
