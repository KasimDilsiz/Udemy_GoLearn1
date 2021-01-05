package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Money penny": 555,
			"Q":           777,
			"M":           888,
		},
		favDrinks: []string{
			"martini",
			"water",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.friends)
	fmt.Println(p1.favDrinks)

	for i, k := range p1.friends {
		fmt.Println(i, k)
	}

	for s, val := range p1.favDrinks {
		fmt.Println(s, val)
	}
}
