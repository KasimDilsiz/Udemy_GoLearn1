package main

import "fmt"

type myType int

var b myType

func main() {

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	b = 42
	fmt.Println(b)
}
