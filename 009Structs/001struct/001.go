package main

import "fmt"

//S001struct , farklı türdeki değerleri bir araya getirmemizi sağlayan bir veri yapısıdır.
//yani toplı bir veri türüdür .
// Ayrıca, farklı türdeki değerleri toplamak için topladığı bir toplu veri türü olduğunu söyleyebiliriz.

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   37,
	}
	p2 := person{
		first: "Kasim",
		last:  "DİLSİZ",
		age:   23,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.last, p1.first)
	fmt.Println(p2.age)
}
