package main

import "fmt"

func main() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	s := fmt.Sprintf("%T\t %T\t %T\t", x, y, z)
	ss := fmt.Sprintf("%v\t %v\t %v\t", x, y, z)
	fmt.Println(s)
	fmt.Println(ss)
}
