package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Starting the Programm")
	fmt.Printf("------------------------\n")
	fmt.Println("Middle of the Programm")
	fmt.Println("------------------------")
	defer fmt.Println("Sum:", add(10, 20))
	fmt.Printf("\n")
	fmt.Printf("------------------------\n")
	fmt.Println("End of the Programm")
	fmt.Printf("------------------------\n")
}
// If defer in more than one statement, then it will execute in LIFO order.(Stack)