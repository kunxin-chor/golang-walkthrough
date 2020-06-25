package main

import "fmt"

type employee struct {
	name           string
	salary         float32
	title          string
	yearsOfService int
}

func createEmployee (name string, salary float32, title string, yearsOfService int) *employee {
	employee := employee{name, salary, title, yearsOfService}
	return &employee
}

// struct are passed by values by default
func printEmployee(e *employee) {	
	fmt.Println("Name:", e.name)
	fmt.Println("Salary:", e.salary)
	fmt.Println("Title:", e.title )
	fmt.Println("Years of Service:", e.yearsOfService)
}

// methods
func (e *employee) printInfo () {
	fmt.Println("Name:", e.name)
	fmt.Println("Salary:", e.salary)
	fmt.Println("Title:", e.title )
	fmt.Println("Years of Service:", e.yearsOfService)
}


func main() {
	// ways of defining a struct
	ceo := employee{"Tan Ah Kow", 3200, "CEO", 3}
	printEmployee(&ceo)

	fmt.Println("------")

	hr := createEmployee("Mary Sue", 4200, "HR Manager", 3)
	hr.printInfo()
}
