package main

import "fmt"

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func Subtract(a, b int) (sub int) {
	sub = a - b
	return
}

func Multiply(a, b int) (mul int) {
	mul = a * b
	return
}

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		err := fmt.Errorf("cannot divide by zero")
		return 0, err
	}
	div := a / b
	return div, nil
}

func main() {
	// Add(10, 20)
	// Subtract(20, 10)
	// Multiply(10, 20)

	
	// ans, _ := Divide(20, 0) // If we don't want to use the error, we can use _
	
	ans, err := Divide(20, 0) 
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)
}