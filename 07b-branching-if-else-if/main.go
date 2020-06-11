package main

import "fmt"

func main() {
	fmt.Println("Score to number grades")
	n := 0
	fmt.Printf("Enter score > ")
	fmt.Scanf("%d", &n)
	if n >= 90 {
		fmt.Println("A+")
	} else if n >= 80 {
		fmt.Println("A")
	} else if n >= 70 {
		fmt.Println("B")
	} else if n >= 60 {
		fmt.Println("C")
	} else if n >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

}
