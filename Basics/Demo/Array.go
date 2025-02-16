package Demo

import "fmt"
func Array(){
	fmt.Print("Declaring an Array in Go\n")
	var arr [5]int  // Declares an array of size 5 with default values (0 for int)
	var arr1 = [3]int{1, 2, 3}  // Declares and initializes an array of size 3 .Uses type inference
	arr2 := [3]int{1, 2, 3}  //shorthand 
	fmt.Println("1st array",arr)
	fmt.Println(arr1, arr2)

	arr2[2] = 20 // Modify an element
	fmt.Println("modified array",arr2) // Output: [1, 2, 20]
	
	// array using for Loop
	fmt.Print("array using for Loop\n")
	array := [4]int{2, 4, 6, 8}  // Declaring an array of size 4 with values

	for i := 0; i < len(array); i++ { // Iterating over the array using a for loop
		fmt.Println(array[i]) // Printing each element of the array
	}
	
	fmt.Println()
	array2 := [3]string{"Go", "is", "awesome"} // Declares an array of size 3 with string elements

	for index, value := range array2 { // Iterates over the array using a for-range loop
		fmt.Println("Index:", index, "Value:", value) // Prints the index and value of each element
	}
}