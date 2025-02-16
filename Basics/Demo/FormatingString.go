package Demo

import "fmt"

func Printingformat(){
	fmt.Println("Printing and Formatting")
 //Print 
 fmt.Print("Hello, ")
 fmt.Print("Gadha Pushpan!") // Output: Hello, Gadha Pushpan!

 //Println
 fmt.Println("Hello,")
 fmt.Println("Gadha Pushpan!") 
 
 //Printf
 	name := "Gadha"
	age := 24

	fmt.Printf("Name: %s, Age: %d\n", name, age)
}