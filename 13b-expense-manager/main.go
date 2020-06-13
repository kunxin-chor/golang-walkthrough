package main

import "fmt"

func main() {
	// slice of expenses
	expenses := []Expense{}
	keepGoing := true
	for keepGoing {
		fmt.Println("Expenses Manager")
		fmt.Println("Choose one:")
		fmt.Println("1. Add expense")
		fmt.Println("2. Display expenses")
		fmt.Println("3. Quit")
		n := 0
		fmt.Scanf("%d", &n)
		
		switch n {
		case 1:
			expenses = addExpense(expenses)
		case 2:
			displayExpense(expenses)
		case 3:
			keepGoing = false
		}
	}
}

func addExpense(expenses []Expense) []Expense {
	fmt.Println("Add expense <title> <category> <cost>")
	expenseName := ""
	category := ""
	var amount float32 = 0.0
	fmt.Scanf("%s %s %f", &expenseName, &category, &amount)

	expense := CreateExpense(expenseName, category, amount)
	expenses = append(expenses, *expense)
	return expenses
	
}

func displayExpense(expenses []Expense) {
	fmt.Printf("Number of expenses: %d\n", len(expenses))
	for _, e := range(expenses) {
		fmt.Println(e)
	}
}
