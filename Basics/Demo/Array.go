package Demo

import "fmt"
func Array(){
	fmt.Print("Declaring an Array in Go\n")
	//var arr1 [5]int  // Declares an array of size 5 with default values (0 for int)
	//var arr2 = [3]int{1, 2, 3}  // Declares and initializes an array of size 3
	arr := [3]int{1, 2, 3}  // Uses type inference

	//fmt.Println(arr1)
	//fmt.Println(arr2)
	fmt.Println("1st array",arr)
	//fmt.Println(arr[0]) // Access first element → 1
	//fmt.Println(arr[1]) // Access second element → 2

	arr[2] = 20 // Modify an element
	fmt.Println("2nd modified array",arr) // Output: [1, 2, 20]
	
	// array using for Loop
	fmt.Print("array using for Loop\n")
	array := [4]int{2, 4, 6, 8}  // Declaring an array of size 4 with values

	for i := 0; i < len(array); i++ { // Iterating over the array using a for loop
		fmt.Println(array[i]) // Printing each element of the array
	}
fmt.Print("\n")
fmt.Println("for each loop")
	array2 := [3]string{"Go", "is", "awesome"} // Declares an array of size 3 with string elements

	for index, value := range array2 { // Iterates over the array using a for-range loop
		fmt.Println("Index:", index, "Value:", value) // Prints the index and value of each element
	}
}