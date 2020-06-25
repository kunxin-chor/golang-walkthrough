package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* This program opens generates 1000,000 numbers in parallel */

func main() {
	noChannels()
	withChannels()
	// withTwoChannels()

}

func noChannels() {
	fmt.Println("No Thread:")
	start := time.Now()
	const upperlimit int = 9000000
	var numbers = make([]int, upperlimit)
	for i := 0; i < upperlimit; i++ {
		numbers[i] = rand.Intn(200)
	}
	fmt.Println("Done")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time passed=", elapsed)
}

func withChannels() {
	
	// create channels
	c1 := make(chan int, 100)
	limit := 9000000

	fmt.Println("One Thread:")
	start := time.Now()	
	go makeNumbers(c1,limit/2)	
	go makeNumbers(c1,limit/2)	

	numbers := make([]int, limit)
	i := 0
	for n:= range c1 {
		numbers[i] = n
		i++
	}
	fmt.Println("Done")
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("Time passed=", elapsed)		
	fmt.Println("Last number generated=", i)
}

func makeNumbers(c chan int, limit int) {
	var upperlimit int = limit
	for i := 0; i < upperlimit; i++ {
		n := rand.Intn(200)
		c <- n
	}
	close(c)
}

// func withTwoChannels() {
// 	fmt.Println("Two threads:")
// 	start := time.Now()
	
// 	// create channels
// 	c1 := make(chan []int)
	
// 	go makeNumbers(c1, 9000000/3)	
// 	go makeNumbers(c1, 9000000/3)
// 	go makeNumbers(c1, 9000000/3)
// 	fmt.Println("Done")

// 	numbers := make([]int, 0)
// 	numbers = append(numbers, <-c1...) 
// 	numbers = append(numbers, <-c1...) 
// 	numbers = append(numbers, <-c1...) 
// 	t := time.Now()
// 	elapsed := t.Sub(start)
// 	fmt.Println("Time passed=", elapsed)		
// 	fmt.Println("Numbers generated=", len(numbers))
// 	// fmt.Println(numbers)
// }

