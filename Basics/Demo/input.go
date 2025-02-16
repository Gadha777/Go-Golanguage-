package Demo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Inputs() {
	reader := bufio.NewReader(os.Stdin)

	// Taking name input
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove trailing newline

	// Taking full name input
	fmt.Print("Enter your full name: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	// Taking age input
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)

	// Printing the final result
	fmt.Printf("Hello, %v! Your full name is %v and you are %v years old.\n", name, fullName, age)
}
