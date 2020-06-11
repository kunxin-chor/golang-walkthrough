package main

import "fmt"

func main() {
	number := 0
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &number)
	if lower, upper := getBounds(10); number > lower && number < upper {
		fmt.Println("number is within bounds")
	} else {
		fmt.Println("number is not within bounds")
	}
}

func getBounds(n int) (int, int) {
	return n - 5, n + 5
}
