package main

import "fmt"

var y = 42

func main() {
	// DECLARED the VARIABLE wih the IDENTIFIER "z"
	// is of the TYPE string
	// and I   ASSIGN the VALUE "shaken , nor stirred"
	//Thıs is a STATIC programming language
	// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
	//not a DYNAMIC programming language

	var z = "Shaken, not stirred"

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	//z = 43
	//fmt.Println(z)
	//fmt.Printf("%T\n", z)

}
