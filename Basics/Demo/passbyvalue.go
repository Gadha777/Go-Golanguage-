package Demo

import "fmt"

func ChangeValue(num int) {
	num = 20 // Changes the value inside the function but does NOT affect the caller
    fmt.Println("Inside changeValue:", num) // Will always print 20 
}
