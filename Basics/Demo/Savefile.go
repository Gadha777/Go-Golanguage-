package Demo
import (
	"fmt"
	"os"
)

func SaveFile(){

	content := []byte("hello Gadha , this is a sample text\n")
	err := os.WriteFile("sample.txt",content, 0644) //0644 set read-write permission

	if err != nil{
		fmt.Println("error file:",err)
	}else{
		fmt.Println("file saved successfully")
	}
}