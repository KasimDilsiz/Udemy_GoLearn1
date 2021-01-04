package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xss := [][]string{xs1, xs2}
	fmt.Println(xss)

	for i, xs := range xss {
		fmt.Println("record: ", i)

		for j, val := range xs {
			fmt.Printf("\t index position: %v\t value: \t %v \n", j, val)
		}

	}

}
