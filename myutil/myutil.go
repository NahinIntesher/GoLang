package myutil

import "fmt"

var Public = "I am a public variable"
var private	= "I am a private variable"
func PrintMessage(message string) { 
	fmt.Println(message)

	fmt.Println(Public) // Public variable can be accessed from any package and should be capitalized
	fmt.Println(private) // Private variable can only be accessed from the same package and should be lowercase

	// Same for the fucntions as well: 
	// For public functions, the first letter should be capitalized
	// And for private functions, the first letter should be lowercase
}