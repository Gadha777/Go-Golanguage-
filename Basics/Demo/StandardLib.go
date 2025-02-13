package Demo

import (
	"fmt"
	"strings"
	"sort"
)

func StandardLib() {
	text := "  Hello Go! Go is great!  "

	fmt.Println(strings.ToUpper(text))       // HELLO GO! GO IS GREAT!
	fmt.Println(strings.ToLower(text))       //   hello go! go is great!  
	fmt.Println(strings.TrimSpace(text))     // Remove spaces
	fmt.Println(strings.Contains(text, "Go")) // Check the text contains "Go" true/false
	fmt.Println(strings.Count(text, "Go"))   // Count occurrences of "Go" Output: 2
	fmt.Println(strings.Replace(text, "Go", "Golang", 1)) // Replace "Go" with "Golang"
	
	names := []string{"Gadha", "Zara", "Alice", "Bob"}
	sort.Strings(names)
	fmt.Println(names)

	//number
	numbers := []int{42, 3, 9, 15, 7}
	sort.Ints(numbers)
	fmt.Println(numbers) // Output: [3 7 9 15 42]
	search:=sort.SearchInts(numbers,9) //Output: we will get the position of the element
	fmt.Println(search)
}