package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var ticket [4]int

	// assign 4 random numbers:
	// for ... range makes a copy of the value
	for i := range ticket {
		ticket[i] = rand.Intn(9) + 1 // number from 0 to 9
	}

	fmt.Print("Winng number is: ")
	for _, number := range ticket {
		fmt.Printf("%d ", number)
	}
	
	fmt.Println("Enter your ticket number, seperated by a space")
	var userTicket [4]int

	fmt.Scan(&userTicket[0], &userTicket[1], &userTicket[2], &userTicket[3])
	printTicket(userTicket[:])


	// get prize
	matches := 0
	for i := range ticket {
		if userTicket[i] == ticket[i] {
			matches++
		}
	}

	switch (matches) {
	case 4:
		fmt.Println("First prize!")
	case 3:
		fmt.Println("Second Prize!")
	case 2:
		fmt.Println("Consolation Prize")
	}

}

func printTicket(ticket []int) {
	for _, number := range ticket {
		fmt.Printf("%d ", number)
	}
	fmt.Println(0)
}