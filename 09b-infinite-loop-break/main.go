package main

import "fmt"

func main() {
	fmt.Println("Enter a number, enter 0 to quit")
	total := 0
	for {
		n := 0
		fmt.Scanf("%d", &n)
		if n > 0 {
			total += n
		} else {
			break
		}

	}
	fmt.Println("Total =",total)
}
