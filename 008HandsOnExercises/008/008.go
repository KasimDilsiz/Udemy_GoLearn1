package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:       {`Shaken, not stirred`, `Martinis`, `Women`},
		`money penny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:            {`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(x)

	for k, v := range x {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}

}
