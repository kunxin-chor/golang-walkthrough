package main

// Expense stores the information for an expenditure
type Expense struct {
	name           string
	category	   string
	amount		   float32
}

// CreateExpense takes in the arguments to create an expense
// returns a pointer to the created object
func CreateExpense (name string, category string, amount float32) *Expense {
	expense := Expense{name, category, amount}
	return &expense
}

