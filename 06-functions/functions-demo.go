package main

import (
	"fmt"
	"math"
)

func main() {
	foobar()
	fmt.Println("Area of circle = ", areaOfCircle(32.0))
	total := addTwo(4,7)
	fmt.Println("4 + 7 =", total)
	higher, lower := doubleAndHalf(4)
	fmt.Println("4 x 2 and 4 / 2 is", higher, "and", lower, "respectively")

}

func foobar() {
	fmt.Println("Foobar!")
}

func areaOfCircle(r float32) float32 {
	return r*r*math.Pi
}

func addTwo(n1, n2 int) int {
	return n1 + n2
}

func doubleAndHalf(n1 float32) (float32,float32) {
	return n1*2, n1/2
}
