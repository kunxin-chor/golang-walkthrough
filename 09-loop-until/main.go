package main

import "fmt"

func main() {
	fmt.Println("Enter a number, enter 0 or less to end")
	total := 0
	n := 0
	for fmt.Scanf("%d", &n); n >= 0; {
		total += n
		fmt.Println("Enter next number")
		fmt.Scanf("%d", &n)
	}
	fmt.Println("Total=", total)
}
