package Demo

import "fmt"
func Loop(){
	//for Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	str := "Hello"
	for i := 0; i < len(str); i++ {
		fmt.Println("Index:", i, "Character:", string(str[i]))
	}
	//break
	fmt.Print("break\n")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break // Stops loop when i == 3
		}
		fmt.Println(i)
	}
	//continues
	fmt.Print("continues\n")

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skips 3 and continues
		}
		fmt.Println(i)
	}

	//Go does not have a separate while loop, but you can achieve the same behavior using for without initialization and increment.
	//while 

	i := 1
	for i <= 5 {
		fmt.Print(i)
		i++ // Manual increment
	}
	fmt.Println()
	
	//range loop 
	arr := [3]string{"Go", "is", "fun"}

	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}
}