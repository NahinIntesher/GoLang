package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("test.txt")

	// if err != nil {
	// 	fmt.Println("Error while creating file:", err)
	// 	return
	// }


	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	for{
		content := make([]byte, 1024)
		_, error := file.Read(content)
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println("Error while reading file:", error)
			return
		}
		fmt.Println(string(content))
	}


	defer fmt.Println("File Closed")
	defer file.Close()

	// content := "Hello World!"
	// _,error := io.WriteString(file, content)
	// if error != nil {
	// 	fmt.Println("Error while writing to file:", error)
	// 	return
	// }
}