package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// fmt.Println("Learning Go. Welcome to the world of Go.")
	// fmt.Println("--------------------------------------------")
	// myutil.PrintMessage("Welcome, Nahin")

	// var a int = 10
	// b := 20
	// fmt.Println(a)
	// fmt.Println(b)
	var a string 
	// fmt.Println(myutil.Public)

	
	fmt.Printf("Enter a name: ")
	// fmt.Scan(&a) // scan until space

	reader := bufio.NewReader(os.Stdin)
	a, _ = reader.ReadString('\n') // _ is used to ignore the error value // scan hole line until enter
	fmt.Printf("Name is: ")
	fmt.Println(a)
	// fmt.Println(myutil.private)

}