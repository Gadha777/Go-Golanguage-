package Demo

import "fmt"
func Map(){
	person := map[string]int{
		"Alice": 25,
		"Bob": 30,
		"Charlie": 35,
	}
	delete(person, "Charlie") // Delete key "Charlie"

	for Name, Age := range person {
		fmt.Println(Name , "-", Age)
	}
	fmt.Println("Total Persons:", len(person)) // Output: Total Persons: 2

	//Nested Maps (Map inside a Map)
	
	employees := map[string]map[string]string{
		"101": {"Name": "Alice", "Department": "HR"},
		"102": {"Name": "Bob", "Department": "IT"},
	}

	fmt.Println(employees["101"]["Name"]) // Output: Alice
}