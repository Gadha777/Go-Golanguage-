package Demo

import "fmt"
func Map(){
	//string as key and int as value
	person := map[string]int{	
		"Alice": 25,
		"Bob": 30,
		"Charlie": 35,
	}
	//looping map
	for Name, Age := range person {
		fmt.Println(Name , "-", Age)
	}
	fmt.Println(person)
	fmt.Println(person["Bob"])
	fmt.Println("Total Persons:", len(person)) // Output: Total Persons: 2

	delete(person, "Charlie") // Delete key "Charlie"

	//Nested Maps (Map inside a Map)
	
	employees := map[string]map[string]string{
		"101": {"Name": "Alice", "Department": "HR"},
		"102": {"Name": "Bob", "Department": "IT"},
	}

	fmt.Println(employees["101"]["Name"]) // Output: Alice
}