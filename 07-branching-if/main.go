package main

import "fmt"

func main() {
	// We use if to decide which code to execute
	n := 3
	if n%3 == 0 {
		fmt.Println("n is a multiple of 3")
	}

	// we can do an initialize before the if
	x :=2; if x % 2 == 0 {
		fmt.Println("x is even")
	}

}
