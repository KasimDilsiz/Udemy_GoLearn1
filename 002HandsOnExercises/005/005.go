package main

import "fmt"

type hotDog int

var o hotDog
var p int
//ıujklş
func main() {

	fmt.Println(o)
	fmt.Printf("%T\n", o)
	o = 42
	p = int(o)

	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T", o)

}
