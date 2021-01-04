package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Kasim"])

	v, ok := m["Kasim"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)

	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}
