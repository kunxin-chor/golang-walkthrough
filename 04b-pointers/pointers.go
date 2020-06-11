package main

import "fmt"

func main() {
	var b int = 42
	var p *int = &b
	fmt.Printf("The address stored in p = %p\n", p)
	fmt.Printf("The address of b = %p\n", &b)
	fmt.Printf("The value referred by p is %d\n", *p)
	fmt.Println("Doing *p += 10")
	*p += 10
	fmt.Printf("Now d = %d\n",b);
}
