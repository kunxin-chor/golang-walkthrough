package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, I am %d years old\n", "Alan", 3)

	cupcakes := 3
	calories := 320.5
	fmt.Printf("One cupcake has %f calories. I ate %d, so in total I consumed %.2f calories", calories, cupcakes, float64(cupcakes)*calories)
}
