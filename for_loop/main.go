package main

import "fmt"

func main() {
	// fmt.Println("Loop 1:")
	// fmt.Println("--------")
	// for i := 0; i < 3; i++ {
	// 	println(i)
	// }

	// counter := 0
	
	// fmt.Println()
	// fmt.Println("Loop 2:")
	// fmt.Println("--------")
	// for {
	// 	println("Infinite loop")
	// 	counter++
	// 	if counter == 3 {
	// 		break
	// 	}
	// }

	nummbers := []int{51, 74, 32, 49, 99}
	for index, value := range nummbers{
		fmt.Println("Index: ", index, "Value: ", value)
	}
}