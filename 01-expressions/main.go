// consider the line below as "black magic" first
package main;

// add in the functionality to show stuff on the console
import "fmt"

// all Go program begins in within this func main block
func main() {

	// just having a single value by itself is legit code but does nothing, and is considered as an error
	// 4

	// having an expression by itself is legit but does nothing, and is considered as an error
	// 4 + 4

	// we can output values and the results of expression using fmt.Println
	fmt.Println(4)
	fmt.Println(3 + 4)
}