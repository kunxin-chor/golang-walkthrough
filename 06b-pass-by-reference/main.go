package main

import "fmt"

func main() {
	// assign integer to 42. Or you can do:
	// var b int = 42
	b := 42
	increaseByOne(&b)
	fmt.Println("b =", b)
}

func increaseByOne(b *int) {
	*b = *b + 1
}
