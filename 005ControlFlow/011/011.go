package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this is not print")
	case (2 == 4):
		fmt.Println("this is not prints 2 ")
	case (4 == 4):
		fmt.Println("gifi")
		fallthrough
	case (5 == 8):
		fmt.Println("not true 1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println(" true 15")

	}
}
