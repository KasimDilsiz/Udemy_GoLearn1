package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   37,
		},
		first: "something cool",
		ltk:   true,
	}

	p2 := person{
		first: "Kasim",
		last:  "DILSIZ",
		age:   23,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.person.first, sa1.person.age, sa1.ltk)
	fmt.Println(p2.age)
}
