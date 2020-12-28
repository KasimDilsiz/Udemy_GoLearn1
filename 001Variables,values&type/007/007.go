package main

import "fmt"

var s int

type hotdog int

var b hotdog

func main() {
	s = 43
	b = 32
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	s = int(b)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

}
