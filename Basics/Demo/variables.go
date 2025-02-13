package Demo

import "fmt"

func Variabledeclaration() {
	fmt.Println("Variable Declaration ")
	//string
	var name1 string = "gadha"
	var name2 = "Gadha"
	name3 := "Gadha Pushpan"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	//Int
	var num1 int = 10  // Explicit type
	var num2 = 20      // Type inference
	num3 := 30         // Short declaration

	fmt.Println(num1, num2, num3)

	//Float
	var f1 float32 = 3.14 // Explicit type
	var f2 = 2.718        // Type inference (defaults to float64)
	f3 := 1.618           // Short declaration

	fmt.Println(f1, f2, f3)
	fmt.Print("\n")
}
