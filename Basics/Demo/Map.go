package Demo

import "fmt"
func Map(){
	person := map[string]int{
		"Alice": 25,
		"Bob": 30}

	for Name, Age := range person {
		fmt.Println(Name , "-", Age)
	}
}