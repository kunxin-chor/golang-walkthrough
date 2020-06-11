package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter your name > ")
	fmt.Scanf("%s", &name)
	fmt.Println("Welcome, "+name+"|")

	fmt.Println("Enter a number and see it increaed by 1")
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(number+1)
}