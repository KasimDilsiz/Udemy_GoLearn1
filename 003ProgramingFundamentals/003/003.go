package main

import (
	"fmt"
)

func main() {
	var x int

	x = 42
	t := 42.23454
	fmt.Println(x)
	fmt.Println(t)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", t)
}
