package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name >")
	name, _ := reader.ReadString('\n')
	fmt.Println("Welcome, " + name)
}
