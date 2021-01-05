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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	for _, b := range m {

		fmt.Println(b.firstName)
		fmt.Println(b.lastName)
		for i, val := range b.favFlavors {
			fmt.Println(i, val)

		}
		fmt.Println("______________")
	}

}
