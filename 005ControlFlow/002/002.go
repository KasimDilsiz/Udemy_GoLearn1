package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("the outer loop%d\n", i)

		for j := 0; j < 5; j++ {
			fmt.Printf(" \t\t The inner loop%d\n", j)
		}
	}
}
