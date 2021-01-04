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
	// son yapttığımızda Kasim adında bir harita oluşturup o haritayı kontrol ettik .dizilerde
	//bir elemanın olup olmadığını böyle bir yolla da takip edebiliriz.

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)

	}

}
