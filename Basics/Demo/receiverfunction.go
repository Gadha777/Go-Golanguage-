package Demo

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Receiver function for Person
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

func Receiverfunction() {
	p1 := Person{Name: "Gadha", Age: 24}
	p1.Greet() // Calling the receiver function
}