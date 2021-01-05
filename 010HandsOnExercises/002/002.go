package main

import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {
	p2 := person{
		firstName: "Kasim",
		lastName:  "DILSIZ",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke"},
	}

	p1 := person{

		firstName: "James",
		lastName:  "Bond",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino"},
	}

	fmt.Println(p1.firstName, p1.lastName)

	for i, k := range p1.favFlavors {
		fmt.Println(i, k)
	}
	fmt.Println(p1.firstName, p1.lastName)
	for i, k := range p2.favFlavors {
		fmt.Println(i, k)
	}

}
