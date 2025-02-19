package main

import (
	"Basics/Demo"
	"fmt"
)

func main() {
	var choice int

	fmt.Println("Select a function to execute:")
	fmt.Println("1. Variable Declaration")
	fmt.Println("2. Standard Library")
	fmt.Println("3. Printing Format")
	fmt.Println("4. Map")
	fmt.Println("5. Array")
	fmt.Println("6. Loop")
	fmt.Println("7. Conditions")
	fmt.Println("8. Functions")
	fmt.Println("9. List")
	fmt.Println("10. Greet")
	fmt.Println("11. Add")
	fmt.Println("12. MultipleReturnValue")
	fmt.Println("13. pass by value")
	fmt.Println("14. Inputs")
	fmt.Println("15. Receiver function")
	fmt.Println("16. Savefile")


	fmt.Print("Enter your choice: ")

	fmt.Scanln(&choice)
	switch choice {
	case 1:
		Demo.Variabledeclaration()
	case 2:
		Demo.StandardLib()
	case 3:
		Demo.Printingformat()
	case 4:
		Demo.Map()
	case 5:
		Demo.Array()
	case 6:
		Demo.Loop()
	case 7:
		Demo.Conditions()
	case 8:
		Demo.Functions("gadha")
	case 9:
		Demo.List([]string{"a", "b", "c"}, Demo.Functions)
	case 10:
		Demo.Greet()
	case 11:
		Demo.Add(5, 3)
	case 12:
		fn, sn :=Demo.MultipleReturnValue("Go Language")
		fmt.Println(fn, sn)
	case 13:
		num1 := 10
		Demo.ChangeValue(num1)
		fmt.Println("Outside function:", num1) // Will still print 10 (unchanged)
	case 14:
		Demo.Inputs()
	case 15:
		Demo.Receiverfunction()
	case 16:
		Demo.SaveFile()
	default:
		fmt.Println("Invalid choice! Please select a valid function.")
	}
}
