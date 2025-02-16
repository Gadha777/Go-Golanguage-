package Demo

import "fmt"

func Greet() {
	fmt.Println("Hello, welcome to Go!")
}

func Add(a int, b int) {
	sum := a + b
	fmt.Println("Sum:", sum)
}
func Functions(n string){
	fmt.Printf("Good morning %v \n",n)
}
func List(n []string, f func (string)){	//passsing arguments as array/slice and function
	for _, value := range n{
		f(value)
	}
}