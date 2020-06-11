package main

import "fmt"

func main() {
	fmt.Println("Variables")

	// declare variable in the format of
	// var <variable name> <variable data type>
	var x int
	var y int

	// declare on the same line
	var a, b int

	// assign using =
	// can only be used for variables that have already been declared
	x = 3
	y = 4

	a = 2
	b = 1
	fmt.Printf("%d", x+y+a+b)

	// init and declare variables at the same time
	var pet1 string = "Fluffy"

	fmt.Println("My dog's name is " + pet1)

	// init and declare multiple variables at the same time

	var student1, student2, student3 string = "Alice", "Bob", "Cindy"

	fmt.Printf("Have you seen %s, %s or %s?\n", student1, student2, student3)

	// SHORTCUT: Declare and assign at the same time
	// ONLY FOR NEW VARIABLES
	pi := 3.14
	radius := 2.5
	fmt.Println("Area of circle is", pi*radius*radius)

}
