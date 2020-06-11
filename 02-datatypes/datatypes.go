package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This is a string")
	fmt.Println(42)   // integer
	fmt.Println(3.14) // float
	fmt.Println(true) // boolean

	// Data type converison
	fmt.Println(int(42.0)) // float to integer
	converted, _ := strconv.ParseInt("42", 10, 32) 
	fmt.Println(converted) // string is in base 10, 32-bit number

	// Values to string
	// first argument is the intger to convert
	// second argument is which base
	fmt.Println(strconv.FormatInt(919, 10))

}
