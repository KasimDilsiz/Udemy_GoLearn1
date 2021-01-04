package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       {`Shaken, not stirred`, `Martinis`, `Women`},
		`money penny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:            {`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(x)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}

}
