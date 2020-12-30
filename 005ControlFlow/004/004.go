package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 8 {
			break
		}
		fmt.Println("`Hello` B1")
		x++

	}
	fmt.Println("\t done")
}
