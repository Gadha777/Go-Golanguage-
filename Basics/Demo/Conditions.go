package Demo

import "fmt"
func Conditions(){
	num := 9

	if num > 10 {
		fmt.Println("Number is greater than 10")
	} else if num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is equal to 10")
	}

	fmt.Println()
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is near")
	case "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("A regular day")
	}
}

//Go allows short variable declaration inside if.
//age is only available inside the if-else block.

// if age := 18; age >= 18 {
// 	fmt.Println("You are an adult")
// } else {
// 	fmt.Println("You are a minor")
// }

//switch with Multiple Cases

// grade := "A"
// 	switch grade {
// 	case "A", "B":
// 		fmt.Println("Good job!")
// 	case "C", "D":
// 		fmt.Println("You can do better")
// 	default:
// 		fmt.Println("Invalid grade")
// 	}

//switch Without a Variable
	// num := 15
	// switch {
	// case num%2 == 0:
	// 	fmt.Println("Even number")
	// case num%2 != 0:
	// 	fmt.Println("Odd number")
	// }