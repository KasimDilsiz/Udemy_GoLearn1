package main

import "fmt"

// DECLARE the variable "y"
//ASSING the value 43
// declare & assing = intiization
var l = 43

//DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
//ASSINGS the ZERO VALUE of TYPE into to "z"
// false for booleans , 0 for integers , 0.0 for floats , ""  for string ,
//and nil for pointers, functions, interfaces ,slices, channels, and maps.

var z int

func main() {
	// short decleration operator
	//DECLARE a variable and ASSING a VALUE( of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(l)
	foo()
	fmt.Println(z)
}
func foo() {
	fmt.Println("again ", l)
}
