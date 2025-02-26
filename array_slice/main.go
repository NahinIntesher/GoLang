package main

import "fmt"

func main() {
	// var a []int = []int{1, 2, 3, 4, 5} // array static
	var a []int = make([]int, 5, 10) // slice dynamic
	var b []string = []string{} // slice dynamic
	
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println(cap(b))
	fmt.Println(len(b))
}
