package main

import "fmt"

func main() {
	
	// form 1
	fmt.Println("Form 1")
	n := 0
	for n < 10 {
		fmt.Println(n)
		n++
	}


	// Form 2
	fmt.Println("Form 2")
	// print to 10
	for x := 0; x < 10; {
		fmt.Println(x)
		x++
	}

	// Form 3
	fmt.Println("Form 3")
	for x:=0; x<10; x++ {
		fmt.Println(x)
	}
}
