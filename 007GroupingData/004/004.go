package main

import "fmt"

func main() {
	c := []int{4, 5, 7, 8, 42}
	fmt.Println(len(c))
	fmt.Println(c[0])
	fmt.Println(c[1])
	fmt.Println(c[2])
	fmt.Println(c[3])
	fmt.Println(c[4])

	for i, v := range c {
		fmt.Println(i, v)
	}

}
