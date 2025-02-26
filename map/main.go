package main

import "fmt"

func main() {

	// name <-> grade
	// studentGrades1 := make(map[string]int)
	// studentGrades1["John"] = 90
	// studentGrades1["Jhosss"] = 90
	// studentGrades1["Nahin"] = 100
	// studentGrades1["Ador"] = 95
	// studentGrades1["Rahim"] = 85
	// studentGrades1["Karim"] = 75
	// studentGrades1["Ananta"] = 80
	// studentGrades1["Jayed"] = 70


	studentGrades1 := map[string]int{
		"John": 90,
		"Jhosss": 90,
		"Nahin": 100,
		"Ador": 95,
		"Rahim": 85,
		"Karim": 75,
		"Ananta": 80,
		"Jayed": 70,
	}


	// Delete a key
	delete(studentGrades1, "Jhosss")

	// Check if a key exists
	// -------------------------------------------
	// grade, exists :=  studentGrades1["Johfn"]
	// fmt.Println("Grade of John is: ", grade)
	// fmt.Println("Exists: ", exists)
	// if exists {
	// 	fmt.Println("Grade of John is: ", grade)
	// } else {
	// 	fmt.Println("John's grade not found")
	// }


	// Iterate over the map 	
	for key, value := range studentGrades1 {
		fmt.Println(key, value)
	}

	
	
	
}